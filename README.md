# learngo

Nomad Coder의 Go 강의 실습용 repo

## bankexample

struct를 이용하여 기본적인 계좌 구조를 만드는 실습

- create
- deposit
- withdraw
- update owner

## dictonaryexample

map을 이용하여 기본적인 사전 구조를 만드는 실습

- search
- create
- update
- delete

## urlchecker

- 여러 개의 URL을 동시에 확인하는 실습
- Go Routine과 Channel을 이용하여 동시에 여러 개의 URL을 확인하는 실습

## jobscrapper

- 웹 페이지에서 채용 공고를 스크래핑하는 실습
  - 실제 강의에서는 indeed를 사용하였으나, 해당 페이지가 변경되어 saramin에서 진행
- 비동기 처리를 위해 Go Routine과 Channel을 사용
- 파일 입출력을 위해 csv 패키지 사용
- echo를 사용하여 웹 서버 구축

---

### To Do

- jobscrapper의 `writeJobs` 함수를 Go Routine을 이용하여 빠르게 동작하도록 하기
