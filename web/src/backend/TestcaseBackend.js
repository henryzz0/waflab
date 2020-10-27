import * as Setting from "../Setting";

export function getTestcases() {
  return fetch(`${Setting.ServerUrl}/api/get-testcases`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}

export function getTestcase(id) {
  return fetch(`${Setting.ServerUrl}/api/get-testcase?id=${id}`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}

export function updateTestcase(id, testcase) {
  return fetch(`${Setting.ServerUrl}/api/update-testcase?id=${id}`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testcase),
  }).then(res => res.json());
}

export function addTestcase(testcase) {
  return fetch(`${Setting.ServerUrl}/api/add-testcase`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testcase),
  }).then(res => res.json());
}

export function deleteTestcase(testcase) {
  return fetch(`${Setting.ServerUrl}/api/delete-testcase`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testcase),
  }).then(res => res.json());
}
