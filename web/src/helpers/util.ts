import dayjs from "dayjs";

import { sha256 } from "./crypto";

const hash = "CyberTect";
const oneHourMS = 3600 * 1000;
const oneDayMS = 24 * oneHourMS;

export function generatePassword(pass: string): string {
  return sha256(hash + sha256(pass + hash));
}

// formatDate 格式化日期
export function formatDate(str: string): string {
  if (!str) {
    return "--";
  }

  return dayjs(str).format("YYYY-MM-DD HH:mm:ss");
}

// isAllowedUser 判断是否允许该用户
export function isAllowedUser(
  allowList: string[],
  currentList: string[]
): boolean {
  if (!allowList || allowList.length === 0) {
    return true;
  }
  let allowed = false;
  allowList.forEach((item) => {
    currentList.forEach((current) => {
      if (current === item) {
        allowed = true;
      }
    });
  });
  return allowed;
}

// today 获取当天0点时间
export function today(): Date {
  return new Date(new Date(new Date().toLocaleDateString()).getTime());
}

// tomorrow 获取明天0点时间
export function tomorrow(): Date {
  return new Date(today().getTime() + oneDayMS);
}

// getDaysAgo 获取多少天前
export function getDaysAgo(days: number): Date {
  return new Date(today().getTime() - days * oneDayMS);
}

// getHoursAge 获取多少小时前
export function getHoursAge(hours: number): Date {
  return new Date(Date.now() - hours * oneHourMS);
}

// today 获取当天0点时间
export function yesterday(): Date {
  return getDaysAgo(1);
}

// formatDateWithTZ 格式化日期（带时区）
export function formatDateWithTZ(date: Date): string {
  return dayjs(date).format("YYYY-MM-DDTHH:mm:ssZ");
}
export function formatBegin(begin: Date): string {
  return formatDateWithTZ(begin);
}
export function formatEnd(end: Date): string {
  return formatDateWithTZ(new Date(end.getTime() + 24 * 3600 * 1000 - 1));
}

interface Shortcut {
  text: string;
  value: Date[];
}

// getDateDayShortcuts 获取时间快捷选择，此返回的时间只到天，在处理的时候应该处理为开始时间的00:00，结束时间的23:59
export function getDateDayShortcuts(ranges: string[]): Shortcut[] {
  const shortcuts: Shortcut[] = [];
  ranges.forEach((element) => {
    switch (element) {
      case "1d":
        shortcuts.push({
          text: "今天",
          value: [today(), today()],
        });
        break;
      case "2d":
        shortcuts.push({
          text: "最近2天",
          value: [getDaysAgo(1), today()],
        });
        break;
      case "3d":
        shortcuts.push({
          text: "最近3天",
          value: [getDaysAgo(2), today()],
        });
        break;
      case "7d":
        shortcuts.push({
          text: "最近7天",
          value: [getDaysAgo(6), today()],
        });
        break;
      default:
        break;
    }
  });
  return shortcuts;
}

// getDateTimeShortcuts 获取时间快捷选择
export function getDateTimeShortcuts(ranges: string[]): Shortcut[] {
  const shortcuts: Shortcut[] = [];
  ranges.forEach((element) => {
    switch (element) {
      case "1h":
        shortcuts.push({
          text: "最近1小时",
          value: [getHoursAge(1), new Date()],
        });
        break;
      case "2h":
        shortcuts.push({
          text: "最近2小时",
          value: [getHoursAge(2), new Date()],
        });
        break;
      case "3h":
        shortcuts.push({
          text: "最近3小时",
          value: [getHoursAge(3), new Date()],
        });
        break;
      case "12h":
        shortcuts.push({
          text: "最近12小时",
          value: [getHoursAge(12), new Date()],
        });
        break;
      case "1d":
        shortcuts.push({
          text: "今天",
          value: [today(), new Date()],
        });
        break;
      default:
        break;
    }
  });
  return shortcuts;
}

function isEqual(value: any, originalValue: any): boolean {
  // 使用json stringify对比是否相同
  return JSON.stringify(value) == JSON.stringify(originalValue);
}

interface DiffInfo {
  modifiedCount: number;
  data: any;
}
// diff  对比两个object的差异
// eslint-disable-next-line
export function diff(current: any, original: any): DiffInfo {
  const data: any = {};
  let modifiedCount = 0;
  Object.keys(current).forEach((key) => {
    const value = current[key];
    if (!isEqual(value, original[key])) {
      data[key] = value;
      modifiedCount++;
    }
  });
  return {
    modifiedCount,
    data,
  };
}

// validateForm validate form
// eslint-disable-next-line
export function validateForm(form: any) {
  return new Promise<void>((resolve, reject) => {
    form.validate((valid: any, rules: any) => {
      if (valid) {
        return resolve();
      }
      const messagesArr: string[] = [];
      Object.keys(rules).forEach((key) => {
        const arr = rules[key];
        arr.forEach((item: any) => {
          messagesArr.push(item.message);
        });
      });
      return reject(new Error(messagesArr.join("，")));
    });
  });
}

// omitNil omit nil(undefined null)
// eslint-disable-next-line
export function omitNil(data: any): any {
  const params: any = {};
  Object.keys(data).forEach((key) => {
    const value = data[key];
    if (value !== undefined && value !== null) {
      params[key] = value;
    }
  });
  return params;
}

// getFieldRules get field rules
// eslint-disable-next-line
export function getFieldRules(fields: any) {
  const rules: any = {};
  fields.forEach((field: any) => {
    if (field.rules) {
      rules[field.key] = field.rules;
    }
  });
  return rules;
}

// getDetectorStatusList get detector status list
export function getDetectorStatusList(): any[] {
  return [
    {
      name: "启用",
      value: 1,
    },
    {
      name: "禁用",
      value: 2,
    },
  ];
}
