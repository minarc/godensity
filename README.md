# godensity
This repository implements DOM-based Content Extraction via Text Density in Go. The project is particularly useful for extracting main content from web pages by analyzing text density, and it has been thoroughly tested on Korean web pages.

📄 DOM-based Content Extraction via Text Density 논문을 기반으로 Go 언어로 구현한 프로젝트입니다. 주로 한국어 웹 페이지들을 대상으로 테스트하였으며, 텍스트 밀도를 분석하여 주요 콘텐츠를 추출하는 데 유용합니다.

![image](https://user-images.githubusercontent.com/11865340/121805896-3d3da200-cc88-11eb-96c2-7468cc94ae78.png)


# Features
Extracts main content from web pages by removing unnecessary elements (e.g., ads, navigation bars, sidebars).
Analyzes text density to determine relevant content blocks.
Efficient DOM traversal using goquery for HTML parsing and manipulation.

# 주요 기능:
광고, 네비게이션 바, 사이드바 등 불필요한 요소를 제거하여 웹 페이지의 주요 콘텐츠를 추출합니다.
텍스트 밀도를 분석하여 관련된 콘텐츠 블록을 식별합니다.
goquery를 활용하여 효율적인 DOM 순회 및 HTML 파싱을 지원합니다.

# How to run?
``` shell
# Clone the repository
gh repo clone minarc/godensity

# Change to the project directory
cd godensity

# Run tests
go test -v .
```
<img width="1124" alt="image" src="https://user-images.githubusercontent.com/11865340/121860599-f73e1800-cd33-11eb-9927-612df92590ef.png">

# Contribution
Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvement. Contributions are always welcome!

기여는 언제나 환영합니다! 버그를 발견했거나 개선 사항이 있다면 자유롭게 이슈를 등록하거나 PR을 제출해주세요.

