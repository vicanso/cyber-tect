import{g as c,E as m}from"./ExConfigEditor.eed62924.js";import{E as n}from"./ExLoading.befd655b.js";import{s as f}from"./index.14cddf8e.js";import{C as s,c as d}from"./configs.45777d5a.js";import{h as p,r as i,G as a}from"./ui.c327b114.js";import{u}from"./naive.650dd42a.js";import"./ExForm.141618cc.js";import"./detector.f4c77658.js";import"./ExFormInterface.d1e14f35.js";import"./common.c9fa31ee.js";var I=p({name:"MockTime",setup(){const r=u(),o=i(0),e=i(!0);return(async()=>{e.value=!0;try{const t=await d();t.id&&(o.value=t.id)}catch(t){f(r,t)}finally{e.value=!1}})(),{id:o,processing:e}},render(){const{id:r,processing:o}=this;if(o)return a(n,null,null);const e=c({category:s.MockTime,name:s.MockTime});return e.push({name:"\u65F6\u95F4\u914D\u7F6E\uFF1A",key:"data",type:"datetime",placeholder:"\u8BF7\u9009\u62E9\u8981Mock\u7684\u65F6\u95F4"}),a(m,{id:r,title:"\u6DFB\u52A0/\u66F4\u65B0MockTime\u914D\u7F6E",description:"\u9488\u5BF9\u5E94\u7528\u65F6\u95F4Mock\uFF0C\u7528\u4E8E\u6D4B\u8BD5\u73AF\u5883\u4E2D\u8C03\u6574\u5E94\u7528\u65F6\u95F4",formItems:e},null)}});export{I as default};