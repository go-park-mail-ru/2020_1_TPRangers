// export default async function ajax(method, url, data = null, callback) {
//     try {
//         let response;
//         if (method === 'POST' && data) {
//             response = await fetch(url, {
//                 method: 'POST',
//                 body: JSON.stringify(data),
//                 headers: {
//                     'Content-Type': 'application/json'
//                 },
//                 credentials: 'include'
//             });
//         } else {
//             response = await fetch(url, {
//                 method: 'GET',
//                 credentials: 'include',
//             });
//         }
//         console.log(response.status);
//         console.log(response.statusText);
//         console.log(response.json);
//         callback(response);
//     } catch (error) {
//         console.error('Ошибка:', error);
//     }
// }

export default function ajax(method, url, body = null, callback) {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url, true);
    xhr.withCredentials = true;

    xhr.addEventListener('readystatechange', function() {
        if (xhr.readyState !== 4) return;

        callback(xhr.status, xhr.responseText);
    });

    if (body) {
        xhr.setRequestHeader('Content-type', 'application/json; charset=utf8');
        xhr.send(JSON.stringify(body));
        return;
    }

    xhr.send();
}
