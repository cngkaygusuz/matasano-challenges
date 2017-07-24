package repeating_otp


type Cipher struct {
	Plaintext []byte
	Key []byte
	Ciphertext []byte
}

func (c *Cipher) Encrypt() {
	if len(c.Key) == 0 {
		panic("length of key cannot be empty")
	}

	if len(c.Plaintext) == 0 {
		panic("length of plaintext cannot be empty.")
	}

	c.Ciphertext = make([]byte, len(c.Plaintext))

	for i := 0; i < len(c.Plaintext); i++ {
		c.Ciphertext[i] = c.Plaintext[i] ^ c.Key[i % len(c.Key)]
	}
}

func (c *Cipher) Decrypt() {
	if len(c.Key) == 0 {
		panic("length of key cannot be empty")
	}

	if len(c.Ciphertext) == 0 {
		panic("length of plaintext cannot be empty.")
	}

	c.Plaintext = make([]byte, len(c.Ciphertext))

	for i := 0; i < len(c.Ciphertext); i++ {
		c.Plaintext[i] = c.Ciphertext[i] ^ c.Key[i % len(c.Key)]
	}
}