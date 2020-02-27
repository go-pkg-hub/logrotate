package logrotate

import (
    "testing"
)

func TestStringToSize(t *testing.T) {
    if StringToSize("0undefined") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0undefined", 0, StringToSize("0undefined"))
    }
    if StringToSize("0.5undefined") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5undefined", 0, StringToSize("0.5undefined"))
    }
    if StringToSize("1undefined") != 1 {
        t.Errorf("Expecting %q = %v, got: %v", "1undefined", 1, StringToSize("1undefined"))
    }
    if StringToSize("1.5undefined") != 1 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5undefined", 1, StringToSize("1.5undefined"))
    }

    if StringToSize("0") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0", 0, StringToSize("0"))
    }
    if StringToSize("0.5") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5", 0, StringToSize("0.5"))
    }
    if StringToSize("1") != 1 {
        t.Errorf("Expecting %q = %v, got: %v", "1", 1, StringToSize("1"))
    }
    if StringToSize("1.5") != 1 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5", 1, StringToSize("1.5"))
    }

    if StringToSize("0k") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0k", 0, StringToSize("0k"))
    }
    if StringToSize("0.5k") != 512 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5k", 512, StringToSize("0.5k"))
    }
    if StringToSize("1k") != 1024 {
        t.Errorf("Expecting %q = %v, got: %v", "1k", 1024, StringToSize("1k"))
    }
    if StringToSize("1.5k") != 1536 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5k", 1536, StringToSize("1.5k"))
    }

    if StringToSize("0kb") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0kb", 0, StringToSize("0kb"))
    }
    if StringToSize("0.5kb") != 512 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5kb", 512, StringToSize("0.5kb"))
    }
    if StringToSize("1kb") != 1024 {
        t.Errorf("Expecting %q = %v, got: %v", "1kb", 1024, StringToSize("1kb"))
    }
    if StringToSize("1.5kb") != 1536 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5kb", 1536, StringToSize("1.5kb"))
    }

    if StringToSize("0m") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0m", 0, StringToSize("0m"))
    }
    if StringToSize("0.5m") != 524288 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5m", 524288, StringToSize("0.5m"))
    }
    if StringToSize("1m") != 1048576 {
        t.Errorf("Expecting %q = %v, got: %v", "1m", 1048576, StringToSize("1m"))
    }
    if StringToSize("1.5m") != 1572864 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5m", 1572864, StringToSize("1.5m"))
    }

    if StringToSize("0mb") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0mb", 0, StringToSize("0mb"))
    }
    if StringToSize("0.5mb") != 524288 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5mb", 524288, StringToSize("0.5mb"))
    }
    if StringToSize("1mb") != 1048576 {
        t.Errorf("Expecting %q = %v, got: %v", "1mb", 1048576, StringToSize("1mb"))
    }
    if StringToSize("1.5mb") != 1572864 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5mb", 1572864, StringToSize("1.5mb"))
    }

    if StringToSize("0g") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0g", 0, StringToSize("0g"))
    }
    if StringToSize("0.5g") != 536870912 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5g", 536870912, StringToSize("0.5g"))
    }
    if StringToSize("1g") != 1073741824 {
        t.Errorf("Expecting %q = %v, got: %v", "1g", 1073741824, StringToSize("1g"))
    }
    if StringToSize("1.5g") != 1610612736 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5g", 1610612736, StringToSize("1.5g"))
    }

    if StringToSize("0gb") != 0 {
        t.Errorf("Expecting %q = %v, got: %v", "0gb", 0, StringToSize("0gb"))
    }
    if StringToSize("0.5gb") != 536870912 {
        t.Errorf("Expecting %q = %v, got: %v", "0.5gb", 536870912, StringToSize("0.5gb"))
    }
    if StringToSize("1gb") != 1073741824 {
        t.Errorf("Expecting %q = %v, got: %v", "1gb", 1073741824, StringToSize("1gb"))
    }
    if StringToSize("1.5gb") != 1610612736 {
        t.Errorf("Expecting %q = %v, got: %v", "1.5gb", 1610612736, StringToSize("1.5gb"))
    }
}
