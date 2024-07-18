Tipe data database -> golang
    VARCHAR, CHAR -> string
    INT, BIGINT -> int32, int64
    FLOAT, DOUBLE -> float32, float64
    BOOLEAN -> bool
    DATE, TIME, DATETIME, TIMESTAMP -> time.Time

Untuk time.Time secara default akan menjadi []byte
    Bisa dikonversi dulu menjadi string lalu di parsing menjadi time.Time
    Ada cara otomatis dari driver MYSQL dengan parsing ?parseDate=true

Nullable Type: Karena golang tidak support null
    string -> database/sql.NullString
    bool -> database/sql.NullBool
    float64 -> database/sql.NullFloat64
    int32 -> database/sql.NullInt32
    int64 -> database/sql.NullInt64
    time.Time -> database/sql.NullTime