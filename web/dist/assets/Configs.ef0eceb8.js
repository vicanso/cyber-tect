import{E as u}from"./ExConfigTable.391ca234.js";import{f as c}from"./configs.6ceae54f.js";import{s as d}from"./index.1e068ea3.js";import{E as m}from"./ExLoading.befd655b.js";import{u as p,M as l,O as g}from"./naive.650dd42a.js";import{h as V,r as o,G as e}from"./ui.c327b114.js";import"./ExTable.6634af00.js";import"./ExFormInterface.d1e14f35.js";import"./common.c9fa31ee.js";const f="\u5F53\u524D\u751F\u6548\u914D\u7F6E";var N=V({name:"ConfigsView",setup(){const r=p(),n=o("\u6240\u6709\u914D\u7F6E"),t=o(""),a=o(!1);return{fetchValid:async()=>{if(!a.value){a.value=!0;try{const s=await c();t.value=JSON.stringify(s,null,2)}catch(s){d(r,s)}finally{a.value=!1}}},tab:n,currentValid:t,fetchingCurrentValid:a}},render(){const{tab:r,fetchValid:n,currentValid:t,fetchingCurrentValid:a}=this;return e(g,{defaultValue:r,onUpdateValue:i=>{i===f&&n()}},{default:()=>[e(l,{name:"\u6240\u6709\u914D\u7F6E"},{default:()=>[e(u,null,null)]}),e(l,{name:f},{default:()=>[a&&e(m,null,null),!a&&e("pre",null,[t])]})]})}});export{N as default};
