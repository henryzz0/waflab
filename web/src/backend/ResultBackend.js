import * as Setting from "../Setting";

export function getResult(testsetId, testcaseId, type) {
  return fetch(`${Setting.ServerUrl}/api/get-result?testsetId=${testsetId}&testcaseId=${testcaseId}&type=${type}`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}
