document.querySelector("#upload").addEventListener("submit", (e) => {
  e.preventDefault();

  const file = e.target.querySelector("input[type=file]").files[0];
  if (!file) {
    return;
  }

  fetch("/signedurl", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      file_name: file.name,
      content_type: "image/png",
    }),
  })
    .then((result) => result.json())
    .then((body) => {
      return fetch(body.upload_url, {
        method: "PUT",
        body: file,
      });
    })
    .catch(alert);
});
