import{F as o}from"./ExFormInterface-d830d527.js";import{u as m,d as c,e as p,f as D,i as d}from"./detector-090ed643.js";import{c as n}from"./ExTable-dcbd254b.js";import{g as l,n as u,a,D as y}from"./Detector-ad8130de.js";import{d as I,N as f}from"./ui-32540f7e.js";import"./index-aeeb7070.js";import"./common-a7d1cc73.js";import"./naive-094f875f.js";import"./configs-5a9029c7.js";import"./ExForm-cdb0faae.js";import"./ExLoading-572e76eb.js";const L=I({name:"DNSDetector",setup(){const e=m().dnsDetectors;return{findByID:async t=>await c(t),dnsDetectors:e}},render(){const{dnsDetectors:e,findByID:s}=this,t=[{title:"域名",key:"host"},n({key:"ips",title:"IP列表"}),n({key:"servers",title:"DNS服务器"})],r=[{name:"域名",key:"host",placeholder:"请输入要检测的域名",span:8},{type:o.DynamicInput,name:"IP列表：",key:"ips",span:8,placeholder:"请输入对应的IP解析",min:1},{type:o.DynamicInput,name:"DNS服务器：",key:"servers",span:8,placeholder:"请输入DNS服务器",min:1}],i=l({host:u("域名不能为空"),servers:a("DNS服务器列表不能为空"),ips:a("域名解析IP地址列表不能为空")});return f(y,{columns:t,fetch:p,findByID:s,updateByID:D,create:d,title:"DNS检测",description:"指定DNS服务器检测解析IP是否正确",formItems:r,data:e,rules:i},null)}});export{L as default};