package playground

// #cgo pkg-config: libavformat libavcodec libavutil libswresample
// #include <libavformat/avformat.h>
// #include <libavcodec/avcodec.h>
// #include <libavutil/avutil.h>
//
// extern int avRead(void*,uint8_t*,int);
// extern int avSeek(void*,int64_t,int);
import "C"
import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"unsafe"

	"github.com/snabb/httpreaderat"
)

type ioProvider struct {
	mu  sync.Mutex
	url string
	off int64
	r   io.ReaderAt
}

func (a *ioProvider) read(p []byte) (int, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.r == nil {
		req, err := http.NewRequest(http.MethodGet, a.url, nil)
		if err != nil {
			return 0, err
		}

		a.r, err = httpreaderat.New(nil, req, nil)
		if err != nil {
			return 0, err
		}
	}

	n, err := a.r.ReadAt(p, a.off)
	a.off += int64(n)

	return n, err
}

func errorCodeToString(ec C.int) error {
	var buf [128]C.char
	C.av_strerror(ec, &buf[0], C.ulong(len(buf)))
	return errors.New(C.GoString(&buf[0]))
}

type FormatInfo struct {
	URL      string `json:"url"`
	Name     string `json:"name"`
	LongName string `json:"long_name"`
	MimeType string `json:"mime_type"`
}

func ProbeFormat(url string) (FormatInfo, error) {
	const bufSize = 4096
	buf := C.av_malloc(bufSize + C.AV_INPUT_BUFFER_PADDING_SIZE)

	var p runtime.Pinner
	defer p.Unpin()

	iop := &ioProvider{url: url}
	p.Pin(iop)

	ctx := C.avio_alloc_context(
		(*C.uchar)(buf),
		bufSize,
		0,
		unsafe.Pointer(iop),
		(*[0]byte)(C.avRead),
		nil,
		(*[0]byte)(C.avSeek),
	)
	defer C.avio_context_free(&ctx)

	var fmt *C.struct_AVInputFormat

	ec := C.av_probe_input_buffer(ctx, &fmt, nil, nil, 0, 0)
	if ec != 0 {
		return FormatInfo{}, errorCodeToString(ec)
	}

	return FormatInfo{
		URL:      url,
		Name:     C.GoString(fmt.name),
		LongName: C.GoString(fmt.long_name),
		MimeType: C.GoString(fmt.mime_type),
	}, nil
}

//export avRead
func avRead(ptr unsafe.Pointer, buf *C.uint8_t, size C.int) C.int {
	p := (*ioProvider)(ptr)
	b := unsafe.Slice((*byte)(buf), int(size))

	n, err := p.read(b)
	if err != nil {
		var ue *url.Error
		if errors.As(err, &ue) {
			return -C.EFAULT
		}

		return -C.EIO
	}

	return C.int(n)
}

//export avSeek
func avSeek(unsafe.Pointer, C.int64_t, C.int) C.int {
	return -C.EPERM
}
