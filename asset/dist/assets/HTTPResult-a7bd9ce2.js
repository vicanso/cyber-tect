import{u as n,x as d}from"./detector-090ed643.js";import{d as r,c as e}from"./ExTable-dcbd254b.js";import{E as u,n as c}from"./ExDetectorResultTable-2191f62a.js";import{N as p}from"./index-aeeb7070.js";import{d as m,N as t}from"./ui-32540f7e.js";import{M as h,C as y,r as k}from"./naive-094f875f.js";import"./ExFormInterface-d830d527.js";import"./common-a7d1cc73.js";const w="p1777rsr",S=m({name:"HTTPResult",setup(){const s=async i=>{await d(i)};return{httpDetectorResults:n().httpDetectorResults,fetch:s}},render(){const{httpDetectorResults:s,fetch:i}=this,o=[{title:"名称",key:"taskName",width:120},r({title:"结果",key:"result.desc",width:80}),{title:"检测URL",key:"url",width:300},{title:"最大耗时(ms)",key:"maxDuration",width:120},e({title:"失败信息",key:"messages"}),{title:"更新于",key:"updatedAt",width:180,render(l){return p(l.updatedAt)}},{title:"更多",key:"",render(l){const a=[{title:"地址",key:"addr",fixed:"left",width:200},{title:"代理地址",key:"proxy",width:200},r({title:"结果",key:"result.desc"}),{title:"失败信息",key:"message",ellipsis:{tooltip:!0},width:200},e({key:"timeline",title:"耗时(ms)",width:180}),{title:"HTTP协议",key:"protocol",width:100},{title:"TLS",key:"tlsVersion",width:80},{title:"TLS加密",key:"tlsCipherSuite",ellipsis:{tooltip:!0},width:100},e({key:"certificateDNSNames",title:"证书域名",width:120}),e({width:140,key:"certificateExpirationDates",title:"证书有效期"}),{title:"Hash",key:"hash",ellipsis:{tooltip:!0},width:100}];return t(y,{placement:"left-end",trigger:"click"},{default:()=>[t("div",{class:w},[t(h,{scrollX:1500,maxHeight:450,columns:a,data:l.results},null)])],...{trigger:c}})}}];return t(k,{title:"HTTP检测结果"},{default:()=>[t(u,{columns:o,data:s,fetch:i,category:"http"},null)]})}});export{S as default};