const myForm = document.getElementById("myform");
const submitButton = document.getElementById("submit-button");

submitButton.onclick = () => {
  const formData = new FormData(myForm);
  const data = JSON.stringify(Object.fromEntries(formData.entries()));

  console.log(data);

  const options = {
    method: "POST",
    body: data,
  };

  fetch("/form", options)
    .then((response) => {
      if (response.status === 200) {
        alert("保存しました");
      } else {
        alert("保存できませんでした");
      }
      return response.json();
    })
    .then((myjson) => {
      console.log(myjson);
    });
};
