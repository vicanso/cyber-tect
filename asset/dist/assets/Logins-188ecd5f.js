import{E as o}from"./ExTable-dcbd254b.js";import{p as n,q as s,k as r,r as l}from"./index-aeeb7070.js";import{F as a}from"./ExFormInterface-d830d527.js";import{d,E as u,N as c}from"./ui-32540f7e.js";import"./naive-094f875f.js";import"./common-a7d1cc73.js";function p(){return[{title:"账户",key:"account",width:120,ellipsis:{tooltip:!0}},{title:"IP",key:"ip",width:120,ellipsis:{tooltip:!0}},{title:"定位",key:"location",width:150},{title:"运营商",key:"isp",width:80},{title:"Track ID",key:"trackID"},{title:"Session ID",key:"sessionID"},{title:"Forwarded For",key:"xForwardedFor",width:140,ellipsis:{tooltip:!0}},{title:"User Agent",key:"userAgent",width:120,ellipsis:{tooltip:!0}}]}function g(){return[{key:"account",name:"账户：",placeholder:"请输入要筛选的账号"},{key:"begin:end",name:"开始结束时间：",type:a.DateRange,span:12,placeholder:"请选择开始时间:请选择结束时间",defaultValue:[s().toISOString(),new Date().toISOString()]}]}const I=d({name:"LoginStats",setup(){const{logins:e}=r(),i=async t=>l({limit:t.limit,offset:t.offset,account:t.account||"",begin:t.begin||"",end:t.end||"",order:"-updatedAt"});return u(()=>{n()}),{logins:e,fetchLogins:i}},render(){const{logins:e,fetchLogins:i}=this;return c(o,{title:"登录查询",filters:g(),columns:p(),data:e,fetch:i},null)}});export{I as default};
