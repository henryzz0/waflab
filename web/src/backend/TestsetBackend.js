import * as Setting from "../Setting";

export function getTestsets() {
  return fetch(`${Setting.ServerUrl}/api/get-testsets`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}

export function getTestset(id) {
  return fetch(`${Setting.ServerUrl}/api/get-testset?id=${id}`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}

export function updateTestset(id, testset) {
  return fetch(`${Setting.ServerUrl}/api/update-testset?id=${id}`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testset),
  }).then(res => res.json());
}

export function addTestset(testset) {
  return fetch(`${Setting.ServerUrl}/api/add-testset`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testset),
  }).then(res => res.json());
}

export function deleteTestset(testset) {
  return fetch(`${Setting.ServerUrl}/api/delete-testset`, {
    method: 'POST',
    credentials: 'include',
    body: JSON.stringify(testset),
  }).then(res => res.json());
}
