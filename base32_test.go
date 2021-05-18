package base32

import (
    "testing"
)

func TestEncodeString(t *testing.T) {
    cases := []struct {
        Input string
        Want  string
    }{
        {
            "Hello World!",
            "91JpRU3F41bPxvkCcgGg",
        },
        {
            "Voken",
            "ASqPPSBE",
        },
        {
            "Vision Network",
            "AsMq6tBFdrG4VsbMeWqQ4TR",
        },
        {
            "Decentralize the Internet",
            "8hJP6sbeeHS62u39f9jJ0W38cMG4juKMcnS6Vsbm",
        },
    }

    for _, test := range cases {
        t.Run("EncodeString", func(t *testing.T) {
            got := EncodeString(test.Input)

            if got != test.Want {
                t.Errorf("encode `%s`: got `%s`, want `%s`", test.Input, got, test.Want)
            }
        })
    }
}

func TestDecodeString(t *testing.T) {
    cases := []struct {
        Input string
        Want  string
    }{
        {
            "91JpRU3F41bPxvkCcgGg",
            "Hello World!",
        },
        {
            "ASqPPSBE",
            "Voken",
        },
        {
            "AsMq6tBFdrG4VsbMeWqQ4TR",
            "Vision Network",
        },
        {
            "8hJP6sbeeHS62u39f9jJ0W38cMG4juKMcnS6Vsbm",
            "Decentralize the Internet",
        },
    }

    for _, test := range cases {
        t.Run("DecodeString", func(t *testing.T) {
            got, err := DecodeString(test.Input)

            assertNoError(t, err)

            if got != test.Want {
                t.Errorf("encode `%s`: got `%s`, want `%s`", test.Input, got, test.Want)
            }
        })
    }
}

func TestIsBase32String(t *testing.T) {
    cases := []struct {
        Input string
        Want  bool
    }{
        {
            "91JpRU3F41bPxvkCcgGg",
            true,
        },
        {
            "ASqPPSBE",
            true,
        },
        {
            "AsMq6tBFdrG4VsbMeWqQ4TR",
            true,
        },
        {
            "8hJP6sbeeHS62u39f9jJ0W38cMG4juKMcnS6Vsbm",
            true,
        },
        {
            "8hJP6sbeeHS62u39f9jJ0W38cMG4juKMcnS6VsbM",
            false,
        },
    }

    for _, test := range cases {
        t.Run("IsBase32String", func(t *testing.T) {
            got := IsBase32String(test.Input)

            if got != test.Want {
                t.Errorf("`%s`: got `%v`, want `%v`", test.Input, got, test.Want)
            }
        })
    }
}

func TestDecodeErrInvalidChecksum(t *testing.T) {
    cases := []string{
        "91JpRU3F41bPxvkCcgGG",
        "ASqPPSBe",
        "AsMq6tBFdrG4VsbMeWqQ4Tr",
        "8hJP6sbeeHS62u39f9jJ0W38cMG4juKMcnS6VsbM",
    }

    for _, test := range cases {
        t.Run("ErrInvalidChecksum", func(t *testing.T) {
            _, err := DecodeString(test)
            assertError(t, err, ErrInvalidChecksum)
        })
    }
}

func assertNoError(t *testing.T, got error) {
    if got != nil {
        t.Fatal("got an error but didnt want one")
    }
}

func assertError(t *testing.T, got error, want error) {
    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %s, want %s", got, want)
    }
}
