import{u as n,B as i}from"./detector-090ed643.js";import{d as r,c as l}from"./ExTable-dcbd254b.js";import{E as c,n as m}from"./ExDetectorResultTable-2191f62a.js";import{N as d}from"./index-aeeb7070.js";import{d as p,N as e}from"./ui-32540f7e.js";import{M as f,C as k,r as y}from"./naive-094f875f.js";import"./ExFormInterface-d830d527.js";import"./common-a7d1cc73.js";const D="pux6hxu",L=p({name:"DatabaseResult",setup(){const t=async a=>{await i(a)};return{databaseDetectorResults:n().databaseDetectorResults,fetch:t}},render(){const{databaseDetectorResults:t,fetch:a}=this,o=[{title:"名称",key:"taskName"},r({title:"结果",key:"result.desc"}),l({title:"连接串列表",key:"uris"}),{title:"最大耗时(ms)",key:"maxDuration"},l({title:"失败信息",key:"messages"}),{title:"更新于",key:"updatedAt",render(s){return d(s.updatedAt)}},{title:"更多",key:"",render(s){const u=[{title:"连接串",key:"uri"},r({title:"结果",key:"result.desc"}),{title:"耗时(ms)",key:"duration"},{title:"失败信息",key:"message"}];return e(k,{placement:"left-end",trigger:"click"},{default:()=>[e("div",{class:D},[e(f,{columns:u,data:s.results},null)])],...{trigger:m}})}}];return e(y,{title:"Database检测结果"},{default:()=>[e(c,{columns:o,data:t,fetch:a,category:"database"},null)]})}});export{L as default};