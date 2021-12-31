import{q as g,s as v}from"./index.1e068ea3.js";import{a as h,b as D,d as S,e as x}from"./configs.6ceae54f.js";import{E as j}from"./ExForm.d5ddb982.js";import{F as f}from"./ExFormInterface.d1e14f35.js";import{E as V}from"./ExLoading.befd655b.js";import{h as F,r as A,G as c}from"./ui.c327b114.js";import{u as I,r as N,A as O,s as T}from"./naive.650dd42a.js";function H(e){return[{name:"\u540D\u79F0\uFF1A",key:"name",disabled:e.name!=null,defaultValue:e.name,placeholder:"\u8BF7\u8F93\u5165\u914D\u7F6E\u540D\u79F0"},{name:"\u5206\u7C7B\uFF1A",key:"category",disabled:!0,defaultValue:e.category},{name:"\u72B6\u6001\uFF1A",key:"status",type:f.Select,placeholder:"\u8BF7\u9009\u62E9\u914D\u7F6E\u72B6\u6001",options:[{label:"\u542F\u7528",value:h.Enabled},{label:"\u7981\u7528",value:h.Disabled}]},{name:"\u5F00\u59CB\u65F6\u95F4\uFF1A",key:"startedAt",type:f.DateTime,placeholder:"\u8BF7\u9009\u62E9\u914D\u7F6E\u751F\u6548\u5F00\u59CB\u65F6\u95F4"},{name:"\u7ED3\u675F\u65F6\u95F4\uFF1A",key:"endedAt",type:f.DateTime,placeholder:"\u8BF7\u9009\u62E9\u914D\u7F6E\u751F\u6548\u7ED3\u675F\u65F6\u95F4"}]}function B(e){const a="data.",t={};Object.keys(e).forEach(d=>{!d.startsWith(a)||(t[d.substring(a.length)]=e[d])});let n=e.data;return Object.keys(t).length!==0&&(n=JSON.stringify(t)),{name:e.name,status:e.status,category:e.category,startedAt:e.startedAt,endedAt:e.endedAt,data:n,description:e.description}}function q(e,a){const t={};return e.name!=a.name&&(t.name=e.name),e.status!=a.status&&(t.status=e.status),e.category!==a.category&&(t.category=e.category),e.startedAt!==a.startedAt&&(t.startedAt=e.startedAt),e.endedAt!==a.endedAt&&(t.endedAt=e.endedAt),e.data!==a.data&&(t.data=e.data),e.description!==a.description&&(t.description=e.description),t}function p(){}var K=F({name:"ExConfigEditor",props:{id:{type:Number,default:0},title:{type:String,required:!0},description:{type:String,required:!0},formItems:{type:Array,required:!0},onSubmitDone:{type:Function,default:p},onBack:{type:Function,default:p}},setup(e){const a=I(),t=e.id!==0,n=A(!1),d=A({}),l=async s=>{if(n.value)return;const{name:r,category:i,status:m,startedAt:k,endedAt:E}=s;if(!r||!i||!m||!k||!E){g(a,"\u914D\u7F6E\u540D\u79F0\u3001\u5206\u7C7B\u3001\u72B6\u6001\u3001\u5F00\u59CB\u65F6\u95F4\u3001\u7ED3\u675F\u65F6\u95F4\u5747\u4E0D\u80FD\u4E3A\u7A7A");return}try{n.value=!0;const u=B(s);if(t){const b=q(u,d.value);if(Object.keys(b).length===0){g(a,"\u6570\u636E\u672A\u4FEE\u6539\u65E0\u9700\u8981\u66F4\u65B0");return}await D({id:e.id,data:b}),d.value=u}else await S(u);e.onSubmitDone()}catch(u){v(a,u)}finally{n.value=!1}},o=[];return e.formItems.forEach(s=>{o.push(Object.assign({},s))}),o.push({name:"\u914D\u7F6E\u63CF\u8FF0\uFF1A",key:"description",type:f.TextArea,placeholder:"\u8BF7\u8F93\u5165\u914D\u7F6E\u63CF\u8FF0",span:24}),t&&(async()=>{n.value=!0;try{const s=await x(e.id);d.value=s,o.forEach(r=>{if(!!r.key)switch(r.key){case"name":r.defaultValue=s.name;break;case"category":r.defaultValue=s.category;break;case"status":r.defaultValue=s.status;break;case"startedAt":r.defaultValue=s.startedAt;break;case"endedAt":r.defaultValue=s.endedAt;break;case"data":r.defaultValue=s.data;break;case"description":r.defaultValue=s.description;break;default:{const i=r.key.split(".");if(i.length===2&&i[0]==="data")try{r.defaultValue=JSON.parse(s.data)[i[1]]}catch(m){console.error(m)}}break}})}finally{n.value=!1}})(),{currentConfig:d,processing:n,onSubmit:l,items:o}},render(){const{title:e,description:a,id:t,onBack:n}=this.$props,{onSubmit:d,processing:l,items:o,currentConfig:y}=this;return l&&t&&!y.id?c(V,null,null):c(T,{show:l},{default:()=>[c(N,null,{default:()=>[c(O,{title:e,onBack:n==p?void 0:n,subtitle:a},{default:()=>[c(j,{formItems:o,onSubmit:d,submitText:t!==0?"\u66F4\u65B0":"\u6DFB\u52A0"},null)]})]})]})}});export{K as E,H as g};
