import { strip } from "../apps-echo/echo.js";

async function loginResponse(endpoint, email, password, aborter) {
    const response = await fetch(`/${endpoint}?q=${encodeURIComponent(email)}&p=${encodeURIComponent(password)}`, {
        signal: aborter.signal,
    });
    const text = await response.text();

    if (response.ok) {
        return text;
    } else {
        throw new Error(text);
    }
}

function errorMessage(msg) {
    if (!document.getElementById("error")) {
        var div = document.createElement("div");
        div.id = "error"
        div.className = "errormsg";
        div.innerText = msg;
                                    
        document.getElementById("main").appendChild(div);
    } else {
        var div = document.getElementById("error");
        div.innerText = msg;
    }
}

function showPassword() {
    var x = document.getElementById("pass");
    var showButton = document.getElementById('show_hide');
    if (x.type === "password") {
        x.type = "text";
        showButton.innerText = "hide";
    } else {
        x.type = "password";
        showButton.innerText = "show";
    }
}

function main() {
    const pass = document.getElementById('pass');
    const email = document.getElementById('email');
    const enterButton = document.getElementById('enter');
    const showButton = document.getElementById('show_hide');
    const inputs = document.querySelectorAll('input');

    let controller;

    inputs.forEach((input) => {
        input.setAttribute('autocomplete', 'off');
        input.setAttribute('autocorrect', 'off');
        input.setAttribute('autocapitalize', 'off');
        input.setAttribute('spellcheck', false);
    });

    const handleIncorrectAnimation = () => {
        pass.classList.remove("incorrect");
        email.classList.remove("incorrect");
    };

    const maxAttempts = 10;
    var attempts = 0;
    const login = async () => {
        //if (localStorage.getItem("locked") == "true") {
        //    errorMessage("too many login attempts");
        //    return;
        //}
        if (attempts >= maxAttempts) {
            return;
        }
        if (controller !== undefined) {
            controller.abort();
        }

        controller = new AbortController();

        try {
            for (const endpoint of ['login']) {
                email.disabled = true;
                pass.disabled = true;
                const responseText = await loginResponse(endpoint, email.value, pass.value, controller);
                const results = JSON.parse(responseText);

                if (results == null || results === false) {
                    attempts++;
                    var attemptslog = ": attempts left: ".concat(maxAttempts-attempts);
                    pass.classList.add("incorrect");
                    email.classList.add("incorrect");

                    pass.addEventListener("animationend", handleIncorrectAnimation, { once: true });
                    email.addEventListener("animationend", handleIncorrectAnimation, { once: true });
                    
                    if (attempts < 5) {
                        errorMessage("invalid credentials");
                    } else if (attempts >= 5 && attempts < maxAttempts) { 
                        errorMessage("invalid credentials".concat(attemptslog));
                    } else {
                        //localStorage.setItem("locked", "true");
                        errorMessage("too many login attempts");
                        pass.value = "";
                        pass.innerText = "";
                    }
                } else {
                    window.location.assign("../apps-echo/index.html");
                }
            }
            email.disabled = false;
            pass.disabled = false;
        } catch (error) {
            errorMessage("server error: cannot validate")
            console.error(error);
        }
    };

    //pass.addEventListener('keydown', (e) => {
    //    if (e.key == 'Backspace') {
    //        pass.innerText = "";
    //        pass.value = "";
    //    }
    //})

    document.addEventListener('keypress', (e) => {
        if (e.key == 'Enter' && pass.validity.valid && email.validity.valid) {
            login();
        }
    });

    enterButton.addEventListener('click', login);
    showButton.addEventListener('click', showPassword);
}

document.addEventListener('DOMContentLoaded', main);

//window.onload = function() {
//    const email = document.getElementById('email');
//    email.value = localStorage.getItem('email');
//}

