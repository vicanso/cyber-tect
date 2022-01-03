import{E as s}from"./ExTable.2a9f5263.js";import{s as n,K as u}from"./index.14cddf8e.js";import{u as d,c,d as m,E as p,e as h,g as f}from"./ExFluxDetail.87b8dedc.js";import{E as y}from"./ExLoading.befd655b.js";import{F as a}from"./ExFormInterface.d1e14f35.js";import{h as g,b as E,z as x,G as i}from"./ui.c327b114.js";import{u as k}from"./naive.650dd42a.js";import"./common.c9fa31ee.js";function T(){return[{title:"\u8D26\u6237",key:"account",width:100,fixed:"left"},{title:"\u65B9\u6CD5",key:"method",width:80},{title:"\u5206\u7C7B",key:"category",width:120},{title:"\u72B6\u6001\u7801",key:"status",width:80},{title:"\u5F02\u5E38",key:"exception",width:80,render(e){return e.exception?"\u662F":"\u5426"}},{title:"TrackID",key:"tid",width:220},{title:"SessionID",key:"sid",width:220},{title:"IP",key:"ip",width:140},{title:"URI",key:"uri",width:200},{title:"\u51FA\u9519\u4FE1\u606F",key:"error",width:80,ellipsis:{tooltip:!0}},{title:"\u5B8C\u6574\u8BB0\u5F55",key:"httpErrorDetail",width:90,align:"center",render(e){return i(p,{measurement:h,data:e,tagKeys:["method","route","category"]},null)}},{title:"\u65F6\u95F4",key:"createdAt",width:180,fixed:"right"}]}const o=[{label:"\u6240\u6709",value:""}];function b(){return[{key:"account",name:"\u8D26\u6237\uFF1A",placeholder:"\u8BF7\u8F93\u5165\u8981\u7B5B\u9009\u7684\u8D26\u53F7"},{key:"category",name:"\u5206\u7C7B\uFF1A",placeholder:"\u8BF7\u9009\u62E9\u8981\u7B5B\u9009\u7684\u5206\u7C7B",type:a.Select,options:o},{key:"exception",name:"\u662F\u5426\u5F02\u5E38",placeholder:"\u8BF7\u9009\u62E9\u662F\u5426\u7B5B\u9009\u5F02\u5E38\u51FA\u9519",type:a.Select,options:[{label:"\u6240\u6709",value:""},{label:"\u662F",value:"true"},{label:"\u5426",value:"false"}]},{key:"limit",name:"\u67E5\u8BE2\u6570\u91CF\uFF1A",type:a.InputNumber,placeholder:"\u8BF7\u8F93\u5165\u8981\u67E5\u8BE2\u7684\u8BB0\u5F55\u6570\u91CF"},{key:"begin:end",name:"\u5F00\u59CB\u7ED3\u675F\u65F6\u95F4\uFF1A",type:a.DateRange,span:12,defaultValue:[u().toISOString(),new Date().toISOString()]}]}var H=g({name:"HTTPErrorStats",setup(){const e=k(),{httpErrorCategories:t,httpErrors:r}=d();return E(async()=>{try{await c()}catch(l){n(e,l)}}),x(()=>{m()}),{httpErrorCategories:t,httpErrors:r}},render(){const{httpErrors:e,httpErrorCategories:t}=this;return t.processing?i(y,null,null):(o.length===1&&t.items.length!==0&&t.items.forEach(r=>{o.push({label:r,value:r})}),i(s,{disableAutoFetch:!0,hidePagination:!0,title:"HTTP\u54CD\u5E94\u51FA\u9519",filters:b(),columns:T(),data:e,fetch:f},null))}});export{H as default};
