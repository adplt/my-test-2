package response

// JSON200 struct
type JSON200 struct {
	Success        bool        `json:"success" example:"true"`
	HTTPStatusCode int         `json:"-"`
	MessageCode    string      `json:"messagecode" example:"M1"`
	Message        string      `json:"message" example:"Horee ... | Berhasil\nYeay ... | Success"`
	Result         interface{} `json:"results"`
}

// JSON400 struct
type JSON400 struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M50"`
	Message        string `json:"message" example:"Oops ... | Data tidak ditemukan\nOops ... | Data not found"`
}

// JSON403 struct
type JSON403 struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M32"`
	Message        string `json:"message" example:"Silakan dicoba kembali | Data tidak valid\nPlease try again | Data not valid"`
}

// JSON500 struct
type JSON500 struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M37"`
	Message        string `json:"message" example:"Yaahh ... | Transaksi tidak dapat di proses. Silahkan coba kembali.\nOh no ... |Transaction Failed. Please try again, or contact CS : 1500195"`
}

type SuccessEmailVerified struct {
	Success        bool        `json:"success" example:"true"`
	HTTPStatusCode int         `json:"-"`
	MessageCode    string      `json:"messagecode" example:"M79"`
	Message        string      `json:"message" example:"Email Verified!|Your data has secured. Have a nice day!"`
}

type SuccessEmailUpdate struct {
	Success        bool        `json:"success" example:"true"`
	HTTPStatusCode int         `json:"-"`
	MessageCode    string      `json:"messagecode" example:"M80"`
	Message        string      `json:"message" example:"Email Has Been Changed!|Your data has secured. Have a nice day!"`
}

type SuccessReOTP struct {
	Success        bool   `json:"success" example:"true"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M5"`
	Message        string `json:"message" example:"Yeay ... | Re-OTP Success"`
}

type SuccessOTP struct {
	Success        bool   `json:"success" example:"true"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M3"`
	Message        string `json:"message" example:"Horee ... | OTP Berhasil\nYeay ... | OTP Success"`
}

// OTPFailed struct
type OTPFailed struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M4"`
	Message        string `json:"message" example:"Yaahh ... | OTP Gagal\nOh no ... | OTP Failed"`
}

// ReOTPFailed struct
type ReOTPFailed struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M6"`
	Message        string `json:"message" example:"Yaahh ... | Re-OTP Gagal\nOh no ... | Re-OTP Failed"`
}

// OTPLImited struct
type OTPLimited struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M53"`
	Message        string `json:"message" example:"Oops ...|Pembatasan Permintaan OTP\nOops ...|Throttling Request OTP"`
}

// EmailExist struct
type EmailExist struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M15"`
	Message        string `json:"message" example:"Oops ... |Email yang Anda masukkan sudah terdaftar pada sistem Bank.\nOops ... |Email that you have entered is already registered in Bank system."`
}

// Warning struct
type Warning struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M10"`
	Message        string `json:"message" example:"Perhatian | Akun kamu akan diblokir setelah 3 kali salah passcode, kamu harus datang ke cabang kami untuk membenarkannya. Jika kamu ragu silahkan gunakan Lupa Passcode\nHold on | Your account will be blocked after 3 incorrect attempts, you must come to our branch to unblocked it. If you are not sure please use Forgot Passcode"`
}

// Block struct
type Block struct {
	Success        bool   `json:"success" example:"false"`
	HTTPStatusCode int    `json:"-"`
	MessageCode    string `json:"messagecode" example:"M11"`
	Message        string `json:"message" example:"Akun Terblokir|Datang ke cabang kami atau hubungi CS di 1500 188 dan Chat via Whatsapp di 0888 8888 888\nAccount Blocked|Please, come to our branch or call CS on 1500 188 and Chat via Whatsapp at 0888 8888 888"`
}

