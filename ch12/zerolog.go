package main

// zerolog-import
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// zerolog-import

func main() {
	//std_log()
	//log_level()
	log_additional()
}

func std_log() {
	// zerolog-std
	log.Printf("Hello World")
	// {"level":"debug","time":"2020-02-13T22:47:19+09:00","message":"Hello World"}
	// zerolog-std
}

func log_level() {
	// log-level
	// info以上のログのみ出力されるように設定
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// それぞれのレベルで出力
	log.Error().Msg("errorで出力")
	log.Info().Msg("infoで出力")
	log.Debug().Msgf("debugで出力: %v", "ただし出力されない")
	// {"level":"error","time":"2020-08-15T21:41:06+09:00","message":"errorで出力"}
	// {"level":"info","time":"2020-08-15T21:41:06+09:00","message":"infoで出力"}
	// log-level
}

func log_additional() {
	// log-additional
	log.Info().
		Str("app", "awesome-app").
		Int("user_id", 1114).Send()
	// log-additional
}

func logger() {
	ctx := context.Background()
	// configured-logger
	logger := log.With().Int("user_id", 1024).Str("path", "/api/user").Str("method", "post").Logger()
	// context.Contextにロガーを登録
	ctx = logger.WithContext(ctx)

	// context.Contextに設定したロガーを取り出し
	newLogger := zerolog.Ctx(ctx)
	newLogger.Print("debug message")
	// configured-logger
}

// response-wrapper
type responseWriterWrapper struct {
	status       int
	responseSize int
	writer       http.ResponseWriter
	request      *http.Request
	start        time.Time
}

func (r *responseWriterWrapper) Flush() {
	flusher := r.writer.(http.Flusher)
	flusher.Flush()
	r.status = 200
}

func (r responseWriterWrapper) Header() http.Header {
	return r.writer.Header()
}

func (r *responseWriterWrapper) Write(content []byte) (int, error) {
	r.responseSize += len(content)
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.writer.Write(content)
}

func (r *responseWriterWrapper) WriteHeader(statusCode int) {
	r.status = statusCode
	r.writer.WriteHeader(statusCode)
}

var _ http.ResponseWriter = &responseWriterWrapper{}
var _ http.Flusher = &responseWriterWrapper{}

// response-wrapper

// response-zerolog-writer
func (r *responseWriterWrapper) MarshalZerologObject(e *zerolog.Event) {
	e.Str("requestMethod", r.request.Method)
	e.Str("requestUrl", r.request.URL.String())
	e.Int64("requestSize", r.request.ContentLength)
	e.Int("status", r.status)
	e.Int("responseSize", r.responseSize)
	e.Str("referer", r.request.Header.Get("Referer"))
	e.Str("latency", time.Now().Sub(r.start).String())
	e.Bool("cacheHit", r.status == 304)
	forwarded := r.request.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		e.Str("remoteIp", forwarded)
	} else {
		e.Str("remoteIp", r.request.RemoteAddr)
	}
	e.Str("protocol", r.request.Proto)
}

var _ zerolog.LogObjectMarshaler = &responseWriterWrapper{}

// response-zerolog-writer

// http-middleware
func NothingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// http-middleware

// with-logger
type contextKey string

const logKey contextKey = "log"

func WithLogger(logger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceId := r.Header.Get("Trace-Id")
			if traceId == "" {
				traceId = xid.New().String()
			}
			logger = logger.With().Str("trace-id", traceId).Logger()
			ctx := context.WithValue(r.Context(), logKey, logger)

			writer := &responseWriterWrapper{
				writer:  w,
				request: r,
				start:   time.Now(),
			}
			next.ServeHTTP(writer, r.WithContext(ctx))
			logger.Info().Object("httpRequest", writer).Send()
		})
	}
}

// with-logger

// get-logger
func GetLogger(ctx context.Context) zerolog.Logger {
	return ctx.Value(logKey).(zerolog.Logger)
}

// get-logger

func useLogger() {
	// use-log-middleware
	var logger = zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Logger()
	mux := http.NewServeMux()
	mux.Handle("/", WithLogger(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})))
	fmt.Println("start listening at :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	// use-log-middleware
}

func useLoggerWithChi() {
	var logger = zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		Logger()
	// use-log-middleware-chi
	r := chi.NewRouter()
	r.Use(WithLogger(logger))
	// use-log-middleware-chi
	mux.Handle("/", WithLogger(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world")
	})))
	fmt.Println("start listening at :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
