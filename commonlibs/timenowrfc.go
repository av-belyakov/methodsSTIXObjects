package commonlibs

import (
	"fmt"
	"time"
)

// TimeNowRFC3339 возвращает текущее время в формате RFC3339
func TimeNowRFC3339() string {
	return fmt.Sprint(time.Now().Format(time.RFC3339))
}
