import * as Setting from "../Setting";

export function getResult(testsetId, testcaseId) {
  return fetch(`${Setting.ServerUrl}/api/get-result?testsetId=${testsetId}&testcaseId=${testcaseId}`, {
    method: "GET",
    credentials: "include"
  }).then(res => res.json());
}
