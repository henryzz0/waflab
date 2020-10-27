import React from "react";
import {message, Tag} from "antd";

export let ServerUrl = '';

export function initServerUrl() {
  const hostname = window.location.hostname;
  ServerUrl = `http://${hostname}:7070`;
}

export function openLink(link) {
  const w = window.open('about:blank');
  w.location.href = link;
}

export function goToLink(link) {
  window.location.href = link;
}

export function getLink(href, text) {
  return <a target="_blank" href={href}>{text}</a>
}

export function showMessage(type, text) {
  if (type === "") {
    return;
  } else if (type === "success") {
    message.success(text);
  } else if (type === "error") {
    message.error(text);
  }
}

export function deepCopy(obj) {
  return Object.assign({}, obj);
}

export function myParseInt(i) {
  const res = parseInt(i);
  return isNaN(res) ? 0 : res;
}

export function addRow(array, row) {
  return [...array, row];
}

export function prependRow(array, row) {
  return [row, ...array];
}

export function deleteRow(array, i) {
  return [...array.slice(0, i), ...array.slice(i + 1)];
}

export function swapRow(array, i, j) {
  return [...array.slice(0, i), array[j], ...array.slice(i + 1, j), array[i], ...array.slice(j + 1)];
}

export function getFormattedDate(date) {
  date = date.replace('T', ' ');
  date = date.replace('+08:00', ' ');
  return date;
}

export function getTagColor(s) {
  if (s === "GET") {
    return "success";
  } else if (s === "POST") {
    return "processing";
  } else if (s === "PUT") {
    return "warning";
  } else if (s === "DELETE") {
    return "error";
  } else {
    return "default";
  }
}

export function getTags(tag) {
  if (tag === undefined || tag === "") {
    return "(None)";
  }

  let res = [];
  const tags = tag.split(",");
  tags.forEach((tag, i) => {
    res.push(
        <Tag color={getTagColor(tag)}>
          {tag}
        </Tag>
    );
  });
  return res;
}

export function getStatusTag(i) {
  if (i === 200) {
    return (
      <Tag color="success">
        200
      </Tag>
    )
  } else if (i === 403) {
    return (
      <Tag color="error">
        403
      </Tag>
    )
  } else {
    return null;
  }
}
