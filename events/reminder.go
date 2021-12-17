package events

import (
	"fmt"
	"net/http"
	"time"
)

func HowLongUntil(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "It's %s. I'm afraid you'll have to do the math yourself.", time.Now().Format(time.RFC3339))
}
