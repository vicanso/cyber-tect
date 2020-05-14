const idField = {
  key: 'id',
  name: 'ID',
  width: '80'
}
const statusField = {
  key: 'statusDesc',
  name: '状态',
  width: '80'
}
const timeoutField = {
  key: 'timeout',
  name: '超时',
  width: '80'
}
const ownerField = {
  key: 'owner',
  name: '拥有者',
  width: '120'
}
const receiversField = {
  key: 'receivers',
  name: '接收账户列表',
  width: '120'
}
const descriptionField = {
  key: 'description',
  name: '描述',
  width: '80'
}
const updatedAtField = {
  key: 'updatedAtDesc',
  name: '更新于',
  width: '180'
}

function getDetectorStatusField () {
  return {
    label: '状态',
    name: 'status',
    type: 'select',
    required: true,
    span: 5,
    options: [
      {
        value: 1,
        label: '启用'
      },
      {
        value: 2,
        label: '禁用'
      }
    ]
  }
}

function getDetectorTimeoutField (options) {
  return Object.assign({
    label: '超时',
    name: 'timeout',
    span: 7,
    required: true,
    placeholder: '请输入超时限制(ms,s)'
  }, options)
}

function getDetectorReceiversField (options) {
  return Object.assign({
    label: '账户列表',
    name: 'receivers',
    type: 'users',
    required: true,
    placeholder: '请选择要接收告警的账户列表'
  }, options)
}

function getDetectorDescriptionField (options) {
  return Object.assign({
    label: '描述',
    name: 'description',
    span: 24,
    type: 'textarea',
    placeholder: '请输入检测描述'
  }, options)
}

export function getDNSFields () {
  return [
    {
      label: '名称',
      name: 'name',
      required: true,
      span: 8,
      placeholder: '请输入名称'
    },
    {
      label: '域名',
      name: 'hostname',
      required: true,
      span: 8,
      placeholder: '请输入要检测的域名'
    },
    {
      label: 'DNS服务器',
      name: 'server',
      required: true,
      span: 8,
      placeholder: '请输入DNS服务器IP'
    },
    getDetectorStatusField(),
    getDetectorTimeoutField(),
    getDetectorReceiversField(),
    getDetectorDescriptionField()
  ]
}

export function getDNSListFields () {
  return [
    idField,
    statusField,
    {
      key: 'name',
      name: '名称'
    },
    {
      key: 'hostname',
      name: '域名'
    },
    {
      key: 'server',
      name: 'DNS服务器'
    },
    timeoutField,
    ownerField,
    receiversField,
    descriptionField,
    updatedAtField
  ]
}

export function getHTTPFields () {
  return [
    {
      label: '名称',
      name: 'name',
      required: true,
      span: 8,
      placeholder: '请输入名称'
    },
    {
      label: 'URL',
      name: 'url',
      required: true,
      span: 8,
      placeholder: '请输入要检测的网址'
    },
    {
      label: 'IP',
      name: 'ip',
      span: 8,
      placeholder: '请输入要指定的域名IP解析，可选'
    },
    getDetectorStatusField(),
    getDetectorTimeoutField(),
    getDetectorReceiversField(),
    getDetectorDescriptionField()
  ]
}

export function getHTTPListFields () {
  return [
    idField,
    statusField,
    {
      key: 'name',
      name: '名称'
    },
    {
      key: 'url',
      name: '检测地址'
    },
    {
      key: 'ip',
      name: 'IP'
    },
    timeoutField,
    ownerField,
    receiversField,
    descriptionField,
    updatedAtField
  ]
}

export function getTCPFields () {
  return [
    {
      label: '名称',
      name: 'name',
      required: true,
      span: 8,
      placeholder: '请输入名称'
    },
    {
      label: 'IP',
      name: 'ip',
      required: true,
      span: 8,
      placeholder: '请输入要检测的IP地址'
    },
    {
      label: '端口',
      name: 'port',
      inputType: 'number',
      required: true,
      span: 8,
      placeholder: '请输入要检测的端口'
    },
    getDetectorStatusField(),
    getDetectorTimeoutField(),
    {
      label: '网络类型',
      name: 'network',
      placeholder: '请输入网络类型，可选'
    },
    getDetectorReceiversField(),
    getDetectorDescriptionField()
  ]
}

export function getTCPListFields () {
  return [
    idField,
    statusField,
    {
      key: 'name',
      name: '名称'
    },
    {
      key: 'ip',
      name: 'IP'
    },
    {
      key: 'port',
      name: '端口'
    },
    {
      key: 'network',
      name: '网络类型'
    },
    timeoutField,
    ownerField,
    receiversField,
    descriptionField,
    updatedAtField
  ]
}

export function getPingFields () {
  return [
    {
      label: '名称',
      name: 'name',
      required: true,
      span: 8,
      placeholder: '请输入名称'
    },
    {
      label: 'IP',
      name: 'ip',
      required: true,
      span: 8,
      placeholder: '请输入要检测的IP地址'
    },
    {
      label: '网络类型',
      name: 'network',
      span: 8,
      placeholder: '请输入网络类型，可选'
    },
    getDetectorStatusField(),
    getDetectorTimeoutField(),
    getDetectorReceiversField(),
    getDetectorDescriptionField()
  ]
}

export function getPingListFields () {
  return [
    idField,
    statusField,
    {
      key: 'name',
      name: '名称'
    },
    {
      key: 'ip',
      name: 'IP'
    },
    {
      key: 'network',
      name: '网络类型'
    },
    timeoutField,
    ownerField,
    receiversField,
    descriptionField,
    updatedAtField
  ]
}
