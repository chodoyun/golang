package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Go Web App on IIS</title>
    <meta charset="UTF-8">
</head>
<body>
    <h1>안녕하세요! Go 웹 애플리케이션입니다</h1>
    <p>이 페이지는 Go로 개발되어 IIS에서 실행되고 있습니다.</p>
    <p>현재 시간: <span id="time"></span></p>
    
    <script>
        document.getElementById('time').textContent = new Date().toLocaleString('ko-KR');
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "Hello from Go API!", "status": "success"}`)
}

func main() {
	// 모든 환경변수 출력 (디버깅용)
	// fmt.Println("Environment variables:")
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	// 환경 변수에서 포트 가져오기 (기본값: 8080)
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = os.Getenv("PORT") // 대체 환경 변수 확인
	}
	if port == "" {
		port = "8080" // 기본값 설정
		fmt.Println("기본 포트(8080)를 사용합니다.")
	} else {
		fmt.Printf("환경 변수에서 포트 %s를 가져왔습니다.\n", port)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)

	fmt.Printf("서버가 포트 %s에서 실행중입니다...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
