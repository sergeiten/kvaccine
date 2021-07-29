# KakaoTalk Vaccine Reservation

In order to detect user who wants to apply for vaccine we need to get user cookies from 잔여백신 tab. We need to use [proxyman](https://proxyman.io/) application to get user cookies. Follow instruction to proper installation proxyman ssl certificate on devices.

## Getting Organization Detail
```bazaar
curl 'https://vaccine.kakao.com/api/v2/org/org_code/41362390' \
-H 'Referer: https://vaccine.kakao.com/detail/41362390?lat=37.46438436921789&lng=127.13949537559671&from=Map' \
-H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2' \
-H 'Host: vaccine.kakao.com' \
-H 'Accept-Language: en-us' \
-H 'Accept: application/json, text/plain, */*' \
-H 'Connection: keep-alive' \
--cookie '_T_ANO=bJyk0j4HdFIjDKDO6g00+56fzNI5NR8LQo5K0XtwIeRSVStn5WIdaV5ziUg95hjnE5uu1V1muUsIRaa52HDvptm96fTYIHQl2WvmqjrgvfeTJMO1xSs/jq9bDQA5XcAK6wNuFhJsGIWJW000GUseDnu0nodr5YXBoyViE91Gzxsc+GSnHeNTE8cnBsEELo1zDia0hEBfRe6XwxMU9Uqv0V+QOvrHKarzCcQnZ8fvoIVqr9YupxPzWURUg/TT64Cxt3ZMk3P0QdJkSVGlaA5e8YcA8Ys0DMy6APzkL3MfmFLuZP8D7eMDmhzoBKhiP3yKPWWFhPFSDGGtBxTuZ/bwpw==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D' \
--proxy http://localhost:9090
```

## Agreement
```bazaar
curl 'https://vaccine.kakao.com/api/v1/agreement' \
-H 'Host: vaccine.kakao.com' \
-H 'Accept-Language: en-us' \
-H 'Referer: https://vaccine.kakao.com/detail/41362390?lat=37.46438436921789&lng=127.13949537559671&from=Map' \
-H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2' \
-H 'Connection: keep-alive' \
-H 'Accept: application/json, text/plain, */*' \
--cookie '_T_ANO=HCRHgYXopZWi/ZONVZIyMK04zcr/HpRj3GZ1CIbP3CiiabZI0RKY6Z4jno8eqF21RGuzNpNuerwdzLmZg5MqXAjxN9DbIfWaq5dDi+LZcHv2i8b4UxJrvSenqVHxrsuQljj80MakeD13RiVP3xUy9xlwXCyWKqjriOkKmCRN5ypr8eITwsjhZ1emqgZo4uC8674r69iEA/gJr5CICBewcsJPbhjDrFg3wGQD+QPgF84IN/hxSjUtJc0GhqvlsReHdogkmdyySdpgoPBdlIEIvJ5qB/RCg9X56Bt3DB3PIM9EU+BX4I5yWhG+CeYurZGz+csagfU5quepLaHLrdWxJw==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D' \
--proxy http://localhost:9090
```

## User Status
```bazaar
curl 'https://vaccine.kakao.com/api/v1/me/status' \
-H 'Connection: keep-alive' \
-H 'Host: vaccine.kakao.com' \
-H 'Accept-Language: en-us' \
-H 'Accept: application/json, text/plain, */*' \
-H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2' \
-H 'Referer: https://vaccine.kakao.com/detail/41362390?lat=37.46438436921789&lng=127.13949537559671&from=Map' \
--cookie '_T_ANO=CyMtKO2YENofoYP7d6oN3op4KPs6qJOSsOdW+LOZ8+q4iTwGrsPoZI8lXDCYCcuba/cz7mZuIF+k/4PonU+CXukA+gcxJlz7OCmZq8wC7QkpUpyCwDMhpA30RG8bUUu/b0rpnZdzMDYQMu2yVOlrYOCuk1V0AeQT6e4lswZFbofcycM/HGelbrPA62qpGsQ7CgcC9K2L/IGbqXTie6tUbsOqmxmeIFp0BG2yzQVu7Dqvz5UnowHpMrZnx27uFCTovWymeDho70Tp2LGw9Ub7ESM57o2pPhtgQQokajBUB1kqJDMHXBdEhuGloHk1WLHvMwZgKhzru28dsywGlL0yLQ==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D' \
--proxy http://localhost:9090
```

## Reservation
```bazaar
curl 'https://vaccine.kakao.com/api/v1/reservation' \
-X POST \
-H 'Accept: application/json, text/plain, */*' \
-H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2' \
-H 'Content-Type: application/json;charset=utf-8' \
-H 'Connection: keep-alive' \
-H 'Origin: https://vaccine.kakao.com' \
-H 'Host: vaccine.kakao.com' \
-H 'Accept-Language: en-us' \
-H 'Content-Length: 74' \
-H 'Referer: https://vaccine.kakao.com/reservation/41362390?from=Map&code=VEN00013' \
--cookie '_T_ANO=CyMtKO2YENofoYP7d6oN3op4KPs6qJOSsOdW+LOZ8+q4iTwGrsPoZI8lXDCYCcuba/cz7mZuIF+k/4PonU+CXukA+gcxJlz7OCmZq8wC7QkpUpyCwDMhpA30RG8bUUu/b0rpnZdzMDYQMu2yVOlrYOCuk1V0AeQT6e4lswZFbofcycM/HGelbrPA62qpGsQ7CgcC9K2L/IGbqXTie6tUbsOqmxmeIFp0BG2yzQVu7Dqvz5UnowHpMrZnx27uFCTovWymeDho70Tp2LGw9Ub7ESM57o2pPhtgQQokajBUB1kqJDMHXBdEhuGloHk1WLHvMwZgKhzru28dsywGlL0yLQ==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D' \
--proxy http://localhost:9090 \
-d '{"from":"Map","vaccineCode":"VEN00013","orgCode":41362390,"distance":null}'
```

## User Information
```bazaar
curl 'https://vaccine.kakao.com/api/v1/user' \
-H 'User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 KAKAOTALK 9.4.2' \
-H 'Accept-Language: en-us' \
-H 'Referer: https://vaccine.kakao.com/reservation/41362390?from=Map&code=VEN00013' \
-H 'Host: vaccine.kakao.com' \
-H 'Connection: keep-alive' \
-H 'Accept: application/json, text/plain, */*' \
--cookie '_T_ANO=CyMtKO2YENofoYP7d6oN3op4KPs6qJOSsOdW+LOZ8+q4iTwGrsPoZI8lXDCYCcuba/cz7mZuIF+k/4PonU+CXukA+gcxJlz7OCmZq8wC7QkpUpyCwDMhpA30RG8bUUu/b0rpnZdzMDYQMu2yVOlrYOCuk1V0AeQT6e4lswZFbofcycM/HGelbrPA62qpGsQ7CgcC9K2L/IGbqXTie6tUbsOqmxmeIFp0BG2yzQVu7Dqvz5UnowHpMrZnx27uFCTovWymeDho70Tp2LGw9Ub7ESM57o2pPhtgQQokajBUB1kqJDMHXBdEhuGloHk1WLHvMwZgKhzru28dsywGlL0yLQ==; _karmt=3QH8z7mTUgSEBIgN6iW7MUZCMafa9ayagGVgGPgPWhc_fdjBI4Vu9IDcp1-WKS--; _karmtea=1627519100; _kawlt=ImicASA-H_xY0oNKMP5xjnJOK31R6XH9bHfF5Avj36VSZ8IIlHx5hiaoU7e7qUEJWBaGqnEyfDWvfYSffcChXIzKWuW8zOTlgJssw9oi9bf3oEG4QIZbeYmpE9NjFmCD; _kawltea=1627508300; TIARA=naJypWuW1DOrrd5XTNGb7tHCMfN_J8R3q-2dA5fSy9kJHxjPgZa3yY95KMyEcBHQ3MOt-cpRA3mJauktu_zjYRCCZJS8zzLZglX-jU.sJSJmeor8lYfJAr7RtQqFwyIBjjVX2y5wRpgCrFEgj4AHTso-_As6h1.u; _kadu=6Zx5L9Mj_lTTGql__1614917897369; PIF=%2F5io3daq4eMjwvQqMfq6Dw%3D%3D' \
--proxy http://localhost:9090
```
