import{F as o}from"./ExFormInterface-d830d527.js";import{E as p}from"./ExTable-dcbd254b.js";import{s as d,O as m}from"./index-aeeb7070.js";import{C as f,u as h}from"./detector-090ed643.js";import{d as i,I as y,J as g,K as k,f as w,N as l}from"./ui-32540f7e.js";import{u as E,e as x}from"./naive-094f875f.js";const T={xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink",viewBox:"0 0 576 512"},b=k("path",{d:"M288 144a110.94 110.94 0 0 0-31.24 5a55.4 55.4 0 0 1 7.24 27a56 56 0 0 1-56 56a55.4 55.4 0 0 1-27-7.24A111.71 111.71 0 1 0 288 144zm284.52 97.4C518.29 135.59 410.93 64 288 64S57.68 135.64 3.48 241.41a32.35 32.35 0 0 0 0 29.19C57.71 376.41 165.07 448 288 448s230.32-71.64 284.52-177.41a32.35 32.35 0 0 0 0-29.19zM288 400c-98.65 0-189.09-55-237.93-144C98.91 167 189.34 112 288 112s189.09 55 237.93 144C477.1 345 386.66 400 288 400z",fill:"currentColor"},null,-1),_=[b],C=i({name:"EyeRegular",render:function(s,a){return y(),g("svg",T,_)}});const v="s97967t";function F(e){return[{name:"结果：",key:"result",placeholder:"请选择要筛选的结果",type:o.Select,defaultValue:m(e,"result"),options:[{label:"所有",value:""},{label:"成功",value:"1"},{label:"失败",value:"2"}],span:5},{name:"耗时：",key:"duration",placeholder:"请输入过滤时长（大于等于）",type:o.InputDuration,span:7},{name:"检测任务：",key:"filterTasks",placeholder:"请选择要筛选的任务",type:o.MultiSelect,options:[],span:12},{key:"startedAt:endedAt",name:"开始结束时间：",placeholder:"请选择开始时间:请选择结束时间",type:o.DateRange,span:18}]}function $(){return l(x,{class:v},{default:()=>[l(C,null,null)]})}const z=i({name:"ExDetectorResultTable",props:{columns:{type:Array,required:!0},data:{type:Object,required:!0},fetch:{type:Function,required:!0},category:{type:String,required:!0}},setup(e){const s=t=>e.fetch(t),a=h().detectorTasks;return w(async()=>{try{await f(e.category)}catch(t){const r=E();d(r,t)}}),{detectorTasks:a,filterFetch:s}},render(){const{detectorTasks:e}=this,{columns:s,data:a}=this.$props,t=F(this.$route.query);return e.processing||e.items.forEach(r=>{let n=t[0],c=!1;t.forEach(u=>{u.key==="filterTasks"&&(c=!0,n=u)}),!(!c||!n.options)&&n.options.push({label:`${r.name}(id:${r.id})`,value:`${r.id}`})}),l(p,{filters:t,columns:s,data:a,fetch:this.filterFetch},null)}});export{z as E,$ as n};