const form = document.querySelector('.url-form');
const submitBtn = document.querySelector('.url-form button');
const formContainer = document.querySelector('.form-container');



const createUrlBox = (url) => {
    const div = document.createElement("div");
    div.classList.add("url-line");

    div.innerHTML = `
    <input type="text" name="short_url" value="${url}" readonly>
    <button type="button">Copy</button>
  `;

    const button = div.querySelector("button");
    const msg = document.createElement("p");
    msg.classList.add("copy-msg");
    div.appendChild(msg);

    button.addEventListener("click", () => {
        navigator.clipboard.writeText(url)
            .then(() => {
                msg.textContent = "Copied!";
                msg.classList.add("show");

                setTimeout(() => {
                    msg.classList.remove("show");
                }, 2000);
            })
            .catch(err => console.error("Copy failed:", err));
    });

    return div;
};



form.addEventListener('submit', async (e) => {
    e.preventDefault();
    const baseUrl = form.base_url.value;
    const shortUrl = form.short_url.value;
    console.log(baseUrl, shortUrl);
    const res = await fetch('/urls', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            baseUrl,
            shortUrl,
        }),
    })
    if (res.ok) {
        const data = await res.json();
        console.log(data.url);
        const url = createUrlBox(data.url);
        formContainer.appendChild(url);
        form.base_url.value = '';
        form.short_url.value = '';
    } else {
        alert('Error creating URL');
    }
});

