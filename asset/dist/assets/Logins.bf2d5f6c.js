import{E as o}from"./ExTable.2a9f5263.js";import{N as n,K as s,E as r,O as l}from"./index.14cddf8e.js";import{F as a}from"./ExFormInterface.d1e14f35.js";import{h as d,z as c,G as u}from"./ui.c327b114.js";import"./naive.650dd42a.js";import"./common.c9fa31ee.js";function f(){return[{title:"\u8D26\u6237",key:"account",width:120,ellipsis:{tooltip:!0}},{title:"IP",key:"ip",width:120,ellipsis:{tooltip:!0}},{title:"\u5B9A\u4F4D",key:"location",width:150},{title:"\u8FD0\u8425\u5546",key:"isp",width:80},{title:"Track ID",key:"trackID"},{title:"Session ID",key:"sessionID"},{title:"Forwarded For",key:"xForwardedFor",width:140,ellipsis:{tooltip:!0}},{title:"User Agent",key:"userAgent",width:120,ellipsis:{tooltip:!0}}]}function p(){return[{key:"account",name:"\u8D26\u6237\uFF1A",placeholder:"\u8BF7\u8F93\u5165\u8981\u7B5B\u9009\u7684\u8D26\u53F7"},{key:"begin:end",name:"\u5F00\u59CB\u7ED3\u675F\u65F6\u95F4\uFF1A",type:a.DateRange,span:12,placeholder:"\u8BF7\u9009\u62E9\u5F00\u59CB\u65F6\u95F4:\u8BF7\u9009\u62E9\u7ED3\u675F\u65F6\u95F4",defaultValue:[s().toISOString(),new Date().toISOString()]}]}var I=d({name:"LoginStats",setup(){const{logins:t}=r(),i=async e=>l({limit:e.limit,offset:e.offset,account:e.account||"",begin:e.begin||"",end:e.end||"",order:"-updatedAt"});return c(()=>{n()}),{logins:t,fetchLogins:i}},render(){const{logins:t,fetchLogins:i}=this;return u(o,{title:"\u767B\u5F55\u67E5\u8BE2",filters:p(),columns:f(),data:t,fetch:i},null)}});export{I as default};
