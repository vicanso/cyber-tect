import{a4 as d,r as f,q as y,s as p,v}from"./index.14cddf8e.js";import{h as m,r as u,G as t,f as c}from"./ui.c327b114.js";import{u as C,v as w,w as l,y as N,d as h,r as k}from"./naive.650dd42a.js";import"./common.c9fa31ee.js";async function g(e){const a=d.replace(":key",e),{data:s}=await f.get(a);return s}async function F(e){const a=d.replace(":key",e);await f.delete(a)}var x=m({name:"CachesView",setup(){const e=C(),a=u(""),s=u(!1),r=u("");return{processing:s,key:a,fetch:async()=>{if(!a.value){y(e,"\u8BF7\u8F93\u5165\u8981\u67E5\u8BE2\u7684key");return}if(!s.value){s.value=!0;try{r.value="";const n=await g(a.value);try{const o=JSON.parse(n.data);r.value=JSON.stringify(o,null,2)}catch{r.value=n.data}}catch(n){p(e,n)}finally{s.value=!1}}},del:async()=>{if(!a.value){y(e,"\u8BF7\u8F93\u5165\u8981\u5220\u9664\u7684key");return}if(!s.value)try{r.value="",await F(a.value),v(e,"\u5DF2\u6210\u529F\u6E05\u9664\u6570\u636E")}catch(n){p(e,n)}finally{s.value=!1}},cacheData:r}},render(){const e="large",{fetch:a,cacheData:s,del:r}=this;return t(k,{title:"\u7F13\u5B58\u67E5\u8BE2\u4E0E\u6E05\u9664"},{default:()=>[t("p",null,[c("session\u7684\u7F13\u5B58\u683C\u5F0F ss:sessionID")]),t(w,{xGap:24},{default:()=>[t(l,{span:12},{default:()=>[t(N,{placeholder:"\u8BF7\u8F93\u5165\u7F13\u5B58\u7684key",size:e,clearable:!0,onUpdateValue:i=>{this.key=i}},null)]}),t(l,{span:6},{default:()=>[t(h,{class:"widthFull",size:e,onClick:()=>a()},{default:()=>[c("\u67E5\u8BE2")]})]}),t(l,{span:6},{default:()=>[t(h,{class:"widthFull",size:e,onClick:()=>r()},{default:()=>[c("\u6E05\u9664")]})]}),s&&t(l,{span:24},{default:()=>[t("pre",null,[s])]})]})]})}});export{x as default};
