import{F as a}from"./ExFormInterface-d830d527.js";import{u as c,t as n,j as m,k as i,l as p}from"./detector-090ed643.js";import{c as d}from"./ExTable-dcbd254b.js";import{g as u,a as D,D as l}from"./Detector-ad8130de.js";import{d as f,N as y}from"./ui-32540f7e.js";import"./index-aeeb7070.js";import"./common-a7d1cc73.js";import"./naive-094f875f.js";import"./configs-5a9029c7.js";import"./ExForm-cdb0faae.js";import"./ExLoading-572e76eb.js";const g=f({name:"TCPDetector",setup(){const t=c().tcpDetectors;return{findByID:async e=>await n(e),tcpDetectors:t}},render(){const{tcpDetectors:t,findByID:r}=this,e=[d({key:"addrs",title:"检测地址"})],o=[{type:a.DynamicInput,name:"地址列表：",key:"addrs",span:12,placeholder:"请输入需要检测的地址，如(IP:Port)",min:1}],s=u({addrs:D("检测地址列表不能为空")});return y(l,{columns:e,fetch:m,findByID:r,updateByID:i,create:p,title:"TCP检测配置",description:"指定IP与端口，定时检测是否可用",formItems:o,data:t,rules:s},null)}});export{g as default};