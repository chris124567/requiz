function write() {
    const questions = JSON.parse("{{json .Term}}");
    for (let i = 0; i < questions.length; i++) {
        if (questions[i].word === undefined) {
            questions[i].word = "";
        }
        if (questions[i].definition === undefined) {
            questions[i].definition = "";
        }
    }

    let useTerm = true;
    let correct = 0,
        incorrect = 0;
    let remaining = questions.length;

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

        (document.getElementById("correct") as HTMLElement).innerHTML = String(correct);
        (document.getElementById("incorrect") as HTMLElement).innerHTML = String(incorrect);
        (document.getElementById("remaining") as HTMLElement).innerHTML = String(remaining);

        const question = document.getElementById("question");
        if (!question) {
            return;
        } else if (questions.length === 0) {
            question.innerHTML = "Complete!";
            return;
        }

        const questionText = useTerm ? questions[0].word : questions[0].definition;
        question.innerHTML = questionText.replace("\n", "<br>").trim();
        if (!useTerm) {
            if (questions[0]._imageUrl !== undefined) {
                question.innerHTML =
                    `<br><img style="max-width: 15rem; max-height: 15rem;" class="img-fluid"  src="` +
                    questions[0]._imageUrl +
                    `">` +
                    question.innerHTML;
            }
        }
    }

    (document.getElementById("answer") as HTMLElement).addEventListener("submit", (event) => {
        event.preventDefault();

        if (questions.length === 0) {
            return;
        }

        const display = document.getElementById("correction");
        if (!display) {
            return;
        }
        display.innerHTML = "";

        const answerText = document.getElementById("answerText");
        if (!answerText) {
            return;
        }

        if (
            (useTerm ? questions[0].definition : questions[0].word).toLowerCase().trim() ===
            (answerText as HTMLInputElement).value.toLowerCase().trim()
        ) {
            correct++;
            remaining--;
        } else {
            incorrect++;
            questions.push(questions[0]);

            const fragment = document.createDocumentFragment();

            const span = document.createElement("span");
            span.innerHTML = "<b>Correction</b>: ";
            fragment.appendChild(span);

            const diff = Diff.diffChars(
                useTerm ? questions[0].definition : questions[0].word,
                (answerText as HTMLInputElement).value
            );
            diff.forEach((part) => {
                const color = part.added ? "green" : part.removed ? "red" : "grey";
                const span = document.createElement("span");
                span.style.color = color;
                span.appendChild(document.createTextNode(part.value));
                fragment.appendChild(span);
            });
            display.appendChild(fragment);
        }

        questions.shift();
        update();

        (answerText as HTMLInputElement).value = "";
    });

    (document.getElementById("useTerm") as HTMLElement).addEventListener("change", () => {
        useTerm = !useTerm;
        update();
    });

    update();
}

document.addEventListener("DOMContentLoaded", write);
