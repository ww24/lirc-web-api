/**
 * remocon.js
 */

const ENDPOINT = "/api/v1";

fetch(`${ENDPOINT}/signals`).then(res => {
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
  return fetch(`${ENDPOINT}/${item.remote}/${item.name}`, {
    method: "POST"
  }).then(res => {
    console.log(res);
    return res.json();
  }).then(data => {
    console.log(data);
  }).catch(err => {
    console.log(err);
    // TODO: 適切なエラーハンドリング
    alert(err);
  });
}
