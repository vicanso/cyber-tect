import{g as m,E as c}from"./ExConfigEditor-da704c67.js";import{E as n}from"./ExLoading-572e76eb.js";import{s as f}from"./index-aeeb7070.js";import{C as i,c as p}from"./configs-5a9029c7.js";import{d,r as s,N as a}from"./ui-32540f7e.js";import{u as l}from"./naive-094f875f.js";import"./ExForm-cdb0faae.js";import"./detector-090ed643.js";import"./ExFormInterface-d830d527.js";import"./common-a7d1cc73.js";const x=d({name:"MockTime",setup(){const r=l(),o=s(0),e=s(!0);return(async()=>{e.value=!0;try{const t=await p();t.id&&(o.value=t.id)}catch(t){f(r,t)}finally{e.value=!1}})(),{id:o,processing:e}},render(){const{id:r,processing:o}=this;if(o)return a(n,null,null);const e=m({category:i.MockTime,name:i.MockTime});return e.push({name:"时间配置：",key:"data",type:"datetime",placeholder:"请选择要Mock的时间"}),a(c,{id:r,title:"添加/更新MockTime配置",description:"针对应用时间Mock，用于测试环境中调整应用时间",formItems:e},null)}});export{x as default};