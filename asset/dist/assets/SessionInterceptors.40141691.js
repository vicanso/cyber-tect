import{E as n}from"./ExConfigEditorList.e2cedbb1.js";import{F as r}from"./ExFormInterface.d1e14f35.js";import{E as i}from"./ExLoading.befd655b.js";import{a1 as m,s as c,a2 as d}from"./index.14cddf8e.js";import{C as p}from"./configs.45777d5a.js";import{h as l,b as u,G as a}from"./ui.c327b114.js";import{u as f}from"./naive.650dd42a.js";import"./ExConfigEditor.eed62924.js";import"./ExForm.141618cc.js";import"./detector.f4c77658.js";import"./ExConfigTable.3a673a73.js";import"./ExTable.2a9f5263.js";import"./common.c9fa31ee.js";var T=l({name:"SessionInterceptors",setup(){const t=f(),{routers:e}=d();return u(async()=>{try{await m()}catch(s){c(t,s)}}),{routers:e}},render(){const{routers:t}=this;if(t.processing)return a(i,null,null);const e=[];t.items.forEach(o=>{e.includes(o.route)||e.push(o.route)}),e.sort();const s=[{type:r.Blank,name:"",key:""},{name:"\u63D0\u793A\u4FE1\u606F\uFF1A",key:"data.message",placeholder:"\u8BF7\u8F93\u5165\u51FA\u9519\u63D0\u793A\u4FE1\u606F"},{name:"\u5141\u8BB8\u8D26\u53F7\uFF1A",key:"data.allowAccount",placeholder:"\u8BF7\u8F93\u5165\u5141\u8BB8\u8D26\u53F7\uFF0C\u591A\u4E2A\u8D26\u53F7\u4EE5,\u5206\u5272"},{name:"\u5141\u8BB8\u8DEF\u7531\uFF1A",key:"data.allowRoutes",type:r.MultiSelect,placeholder:"\u8BF7\u8F93\u5165\u5141\u8BB8\u8DEF\u7531\uFF0C\u53EF\u4EE5\u591A\u9009",options:e.map(o=>({label:o,value:o}))}];return a(n,{listTitle:"Session\u62E6\u622A\u914D\u7F6E",editorTitle:"\u6DFB\u52A0/\u66F4\u65B0Session\u62E6\u622A\u914D\u7F6E",editorDescription:"\u8BBE\u7F6ESession\u62E6\u622A\u7684\u76F8\u5173\u914D\u7F6E",category:p.SessionInterceptor,extraFormItems:s},null)}});export{T as default};
