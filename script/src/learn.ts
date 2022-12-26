function learn() {
    const questions = JSON.parse("{{json .Term}}");
    const allQuestions = JSON.parse("{{json .Term}}");

    for (let i = 0; i < questions.length; i++) {
        if (questions[i].word === undefined) {
            questions[i].word = "";
        }
        if (questions[i].definition === undefined) {
            questions[i].definition = "";
        }
    }

    let lock = false;
    let useTerm = true;
    let correct = 0,
        incorrect = 0;
    let remaining = questions.length;

    function min(x: number, y: number) {
        return x < y ? x : y;
    }

    function undefinedString(s: string) {
        return s !== undefined ? s : "";
    }

    function shuffle(array: any[]) {
        for (let i = array.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            const temp = array[i] as any;
            array[i] = array[j] as any;
            array[j] = temp;
        }
    }

    function update() {
        shuffle(questions);
        shuffle(allQuestions);

        (document.getElementById("correct") as HTMLElement).innerHTML = String(correct);
        (document.getElementById("incorrect") as HTMLElement).innerHTML = String(incorrect);
        (document.getElementById("remaining") as HTMLElement).innerHTML = String(remaining);

        const question = document.getElementById("question");
        if (!question) {
            return;
        } else if (questions.length === 0) {
            lock = true;
            question.innerHTML = "Complete!";
            return;
        }

        const questionText = useTerm
            ? undefinedString(questions[0].word)
            : undefinedString(questions[0].definition);
        question.innerHTML = questionText.replace("\n", "<br>").trim();
        if (!useTerm && questions[0]._imageUrl !== undefined) {
            question.innerHTML =
                `<br><img style="max-width: 15rem; max-height: 15rem;" src="` +
                questions[0]._imageUrl +
                `"><br>` +
                question.innerHTML;
        }

        const options = [];
        options.push({
            correct: true,
            _imageUrl: questions[0]._imageUrl,
            text: useTerm
                ? undefinedString(questions[0].definition)
                : undefinedString(questions[0].word),
        });

        let i = 0;
        while (options.length < min(4, allQuestions.length - 1) && i < allQuestions.length) {
            const value = {
                correct: false,
                _imageUrl: allQuestions[i]._imageUrl,
                text: useTerm
                    ? undefinedString(allQuestions[i].definition)
                    : undefinedString(allQuestions[i].word),
            };
            let includes = false;
            options.forEach((option) => {
                if (option.text === value.text && option._imageUrl === value._imageUrl) {
                    includes = true;
                }
            });
            if (!includes) {
                options.push(value);
            }
            i++;
        }

        shuffle(options);
        for (let i = 0; i < options.length; i++) {
            const option = options[i];
            if (!option) {
                return;
            }
            const optionElement = document.getElementById("option" + (i + 1));
            if (!optionElement) {
                console.log("no optionElement:", optionElement);
                return;
            }

            optionElement.innerHTML =
                `<div class="card-body">` + option.text.replace("\n", "<br>").trim() + `</div>`;
            if (useTerm && option._imageUrl !== undefined) {
                optionElement.innerHTML +=
                    `<img style="max-width: 15rem; max-height: 15rem;" class="card-img-top" src="` +
                    option._imageUrl +
                    `">` +
                    optionElement.innerHTML;
            }

            if (option.correct) {
                optionElement.setAttribute("correct", "correct");
            } else if (optionElement.getAttribute("correct") === "correct") {
                optionElement.removeAttribute("correct");
            }
        }
    }

    (document.getElementById("useTerm") as HTMLElement).addEventListener("change", () => {
        useTerm = !useTerm;
        update();
    });

    function optionClick(id: string) {
        if (lock) {
            return;
        }

        const option = document.getElementById(id);
        if (!option) {
            return;
        }
        const originalClass = option.getAttribute("class");
        if (option.getAttribute("correct") === "correct") {
            correct++;
            remaining--;
            option.setAttribute("class", originalClass + " list-group-item-success");
        } else {
            incorrect++;
            option.setAttribute("class", originalClass + " list-group-item-danger");
            questions.push(questions[0]);
        }
        lock = true;

        setTimeout(() => {
            questions.shift();
            option.setAttribute("class", originalClass as string);
            update();
            lock = false;
        }, 265);
    }

    (document.getElementById("option1") as HTMLElement).addEventListener("click", () =>
        optionClick("option1")
    );
    (document.getElementById("option2") as HTMLElement).addEventListener("click", () =>
        optionClick("option2")
    );
    (document.getElementById("option3") as HTMLElement).addEventListener("click", () =>
        optionClick("option3")
    );
    (document.getElementById("option4") as HTMLElement).addEventListener("click", () =>
        optionClick("option4")
    );
    document.onkeypress = (event) => {
        switch (event.code) {
            case "Digit1":
                optionClick("option1");
                break;
            case "Digit2":
                optionClick("option2");
                break;
            case "Digit3":
                optionClick("option3");
                break;
            case "Digit4":
                optionClick("option4");
                break;
        }
    };

    update();
}

document.addEventListener("DOMContentLoaded", learn);
