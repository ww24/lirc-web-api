/**
 * remocon.js
 */

const ENDPOINT = "/api/v1";

fetch(`${ENDPOINT}/`).then(res => {
  console.log(res);
  return res.json();
}).then(data => {
  console.log(data);
  riot.mount("signals", {api: data, send: send});
}).catch(err => {
  console.log(err);
  // TODO: 適切なエラーハンドリング
  alert(err);
});

function send(item) {
  console.log(item);
  // TODO: 送信前処理
  fetch(`${ENDPOINT}/${item.remote}/${item.name}`, {
    method: "POST"
  }).then(res => {
    console.log(res);
    return res.json();
  }).then(data => {
    console.log(data);
    // TODO: 送信後処理
    // alert(`sent signal: ${item.remote}:${item.name}`);
  }).catch(err => {
    console.log(err);
    // TODO: 適切なエラーハンドリング
    alert(err);
  })
}
