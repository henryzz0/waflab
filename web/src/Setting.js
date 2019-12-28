
export let ServerUrl = '';

export function initServerUrl() {
  const hostname = window.location.hostname;
  ServerUrl = `http://${hostname}:7070`;
}
