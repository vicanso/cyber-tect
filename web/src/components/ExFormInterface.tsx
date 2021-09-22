export interface FormItem {
  name: string;
  key: string;
  type?: string;
  placeholder?: string;
  span?: number;
  defaultValue?: unknown;
  disabled?: boolean;
  // TODO 确认是否有其它方式表示
  // eslint-disable-next-line
  options?: any[];
  suffix?: string;
  min?: number;
  max?: number;
}

export enum FormItemTypes {
  Select = "select",
  MultiSelect = "multiSelect",
  DateTime = "dateTime",
  DateRange = "dateRange",
  InputNumber = "inputNumber",
  InputNumberGroup = "inputGroup",
  InputDuration = "inputDuration",
  DynamicInput = "dynamicInput",
  TextArea = "textArea",
  Blank = "blank",
  MultiUserSelect = "multiUserSelect",
}
