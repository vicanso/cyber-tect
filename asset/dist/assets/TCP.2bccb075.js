import{F as a}from"./ExFormInterface.d1e14f35.js";import{u as c,t as n,k as m,l as i,m as d}from"./detector.f4c77658.js";import{c as p}from"./ExTable.2a9f5263.js";import{g as f,D as u,a as D}from"./Detector.f510725a.js";import{h as l,G as I}from"./ui.c327b114.js";import"./index.14cddf8e.js";import"./common.c9fa31ee.js";import"./naive.650dd42a.js";import"./configs.45777d5a.js";import"./ExForm.141618cc.js";import"./ExLoading.befd655b.js";var E=l({name:"TCPDetector",setup(){const t=c().tcpDetectors;return{findByID:async e=>await d(e),tcpDetectors:t}},render(){const{tcpDetectors:t,findByID:r}=this,e=[p({key:"addrs",title:"\u68C0\u6D4B\u5730\u5740"})],s=[{type:a.DynamicInput,name:"\u5730\u5740\u5217\u8868\uFF1A",key:"addrs",span:12,placeholder:"\u8BF7\u8F93\u5165\u9700\u8981\u68C0\u6D4B\u7684\u5730\u5740\uFF0C\u5982(IP:Port)",min:1}],o=f({addrs:D("\u68C0\u6D4B\u5730\u5740\u5217\u8868\u4E0D\u80FD\u4E3A\u7A7A")});return I(u,{columns:e,fetch:n,findByID:r,updateByID:m,create:i,title:"TCP\u68C0\u6D4B\u914D\u7F6E",description:"\u6307\u5B9AIP\u4E0E\u7AEF\u53E3\uFF0C\u5B9A\u65F6\u68C0\u6D4B\u662F\u5426\u53EF\u7528",formItems:s,data:t,rules:o},null)}});export{E as default};