# la

Platforma do ładnych ankiet, może być jedna, mogą być dwie, ale nikt nie wie jak je utworzyć!

Docelowo to ma być platforma do ankiet, może być uruchomiona na serwerze specjalnie dla jednej czy kilku ankiet, może być też dla wielu więcej.

## Wymagania
* MongoDB

Jeżeli chcesz samemu skompilować (skompilowaną wersję oferujemy jedynie na linuxa), to także:
* Go
* Node.js

## Jak sprawić, aby działało?
**Najpierw** pobierz z Releases zipa i go gdzieś rozpakuj.

Tworzymy plik config.json i uzupełniamy go jak w config.json.example, szczegółowo wyjaśnione:
### Captcha
Korzystamy z captchy v3, więc należy skonfigurować wszystko dobrze na [stronie recaptchy](https://www.google.com/recaptcha/admin)

Po stworzeniu, dodaniu stron (jeśli uruchamiasz na swoim komputerze dla siebie to tam dodajesz localhost), interesuje nas zakładka **Keys**

* Site key: kopiujemy do *captchaSite* w config.json
* Secret key: kopiujemy do *captchaSecret* w config.json
### Aplikacja Google
no tutaj ja szczerze powiem że google niepotrzebnie skomplikowało ten proces, ale sobie poradzimy

[aktualnie sama ankieta będzie działała bez tego, ale nie będzie można się zalogować więc lepiej to zrobić na przyszłość]
* Wchodzimy na https://console.cloud.google.com
* Jak nie jesteśmy zalogowani, to się logujemy
* U góry klikamy taką strzałeczkę, później *NOWY PROJEKT*
![ilustracja 1](https://i.imgur.com/Iymp7EH.png)
* Nazywamy nasz projekt i klikamy UTWÓRZ
* Google może nasz nie przenieść do nowego projektu, więc klikamy na strzałeczkę (tą samą którą klikaliśmy wcześniej) i wybieramy nasz nowy projekt
* Teraz z bocznego menu najeżdzamy na `Interfejsy API i usługi` i z wysuniętego panelu wybieramy `Dane logowania`
![ilustracja 2](https://i.imgur.com/q1r04in.png)
* Teraz górnego ale nie najbardziej górnego paska wybieramy `Ekran akceptacji OAuth`
* Ustalamy nazwę aplikacji, najlepiej taką samą jak nazwa projektu
* **Nie ustawiaj loga aplikacji**, bo wtedy zły gogiel każe ci zweryfikować aplikacje, a pewnie złodzieje później będą chcieli dane twojej karty kredytowej jak to mają w zwyczaju
* Możesz ustawić email, a teraz przewijasz na dół i naciskasz `Zapisz`
* Widać dosyć niewielki, niebieski przycisk `Utwórz dane logowania`. Jak można się domyśleć, należy w niego kliknąć. Wysunie się mały pasek, wybieramy `ID klienta OAuth`
![ilustracja 3](https://i.imgur.com/ZQazgok.png)
* Jako typ aplikacji wskazujemy Aplikacja internetowa
* Jako nazwę aplikacji ponownie najlepiej wpisać nazwę projektu
* Teraz bardzo ważne. Wiesz pod którym adresem ma być uruchomiony serwer ankiety? Domyślnie jest to `http://localhost:8080`. Pod **Autoryzowane identyfikatory URI przekierowania** wklejamy właśnie ten adres i do tego `/api/callback`. Czyli jak mamy `http://localhost:8080`, to tam wpisujemy `http://localhost:8080/api/callback`
* Teraz `Utwórz`
* **No i mamy to co chcieliśmy**
* Twój identyfikator klienta: wklejamy do googleClientID w config.json
* Twój tajny klucz klienta: wklejamy do googleSecret w config.json
* No i teraz Google, po co tyle roboty? Zwykły Janusz by sobie z tym nie poradził. Ty też, gdybyś był takim Januszem, to byś sobie nie poradził. A istnieje na to lepszy sposób? Jak sobie nie poradziłeś, to jeśli znasz mojego discorda to możesz się zgłosić po pomoc. A jak sobie poradziłeś, to też możesz się zgłosić to dam ci nagrodę.
### To teraz odpalamy
* Tylko upewnij się że masz odpalone MongoDB
* Aby uruchomić, najlepiej odpalić konsolę i tam wpisać `./la`, bo jak odpalisz z menedżera plików to nie będziesz miał konsolki z logami.
* Domyślnie serwer jest uruchomiany na porcie 8080, więc w przeglądarce wystarczy wejść pod adres `http://localhost:8080`
* Możesz utworzyć przykładową ankietę wchodząc pod adres `/real`.
## Czego nie ma a musi być
* Edytowanie ankiety
// * Logowanie bo aktualnie to nic nie robi
// * Dodawanie ankiet (będzie mogło być zablokowane)
// * Sprawdzanie odpowiedzi

## Screen dla zachęty
![pokazane tu jest](https://i.imgur.com/BpBRLP5.png)

## a może zechciałbyś zostać jestem hardkorem, hakerem lepszym niż qookie, i skompilować to samemu?
### Go
* ```go get github.com/labstack/echo/...```
* ```go get golang.org/x/oauth2/...```
* ```go get github.com/globalsign/mgo/...```

### Kompilowanie vue
* Wchodzimy do `/client`
* `npm install`
* `npm run build`

### I teraz
* `cd ..`
* `go run server.go`
* a tylko nie zapomnij o config.json