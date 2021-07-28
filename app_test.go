package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCookies(t *testing.T) {
	cookie := "_T_ANO=CyMtKO2YENofoYP7d6oN3op4KPs6qJOSsOdW+LOZ8+q4iTwGrsPoZI8lXDCYCcuba/cz7mZuIF+k/4PonU+CXukA+gcxJlz7OCmZq8wC7QkpUpyCwDMhpA30RG8bUUu/b0rpnZdzMDYQMu2yVOlrYOCuk1V0AeQT6e4lswZFbofcycM/HGelbrPA62qpGsQ7CgcC9K2L/IGbqXTie6tUbsOqmxmeIFp0BG2yzQVu7Dqvz5UnowHpMrZnx27uFCTovWymeDho70Tp2LGw9Ub7ESM57o2pPhtgQQokajBUB1kqJDMHXBdEhuGloHk1WLHvMwZgKhzru28dsywGlL0yLQ==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D"

	cookies := parseCookies(cookie)

	assert.Nil(t, cookies)
}
