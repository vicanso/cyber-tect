import{r as c,Z as o,_ as a,$ as u}from"./index.1e068ea3.js";import{e as f,a as l}from"./ui.c327b114.js";var s;(function(n){n.MockTime="mockTime",n.BlockIP="blockIP",n.SignedKey="signedKey",n.RouterConcurrency="routerConcurrency",n.SessionInterceptor="sessionInterceptor",n.RequestConcurrency="requestConcurrency",n.Router="router",n.Email="email",n.HTTPServerInterceptor="httpServerInterceptor"})(s||(s={}));var r;(function(n){n[n.Enabled=1]="Enabled",n[n.Disabled=2]="Disabled"})(r||(r={}));const e=f({processing:!1,items:[],count:-1});function d(n){n.key=`${n.id}`}async function y(n){const{data:t}=await c.post(o,n);return t}async function g(){const{data:n}=await c.get(o,{params:{category:s.MockTime,name:s.MockTime,limit:1}}),t=n.configurations||[];return t.length===0?{}:t[0]}async function k(n){const t=a.replace(":id",`${n}`),{data:i}=await c.get(t);return i}async function T(n){const t=a.replace(":id",`${n.id}`);await c.patch(t,n.data)}async function b(n){if(!e.processing){n.limit||(n.limit=50);try{e.processing=!0;const{data:t}=await c.get(o,{params:n}),i=t.count||0;i>=0&&(e.count=i),e.items=t.configurations||[],e.items.forEach(d)}finally{e.processing=!1}}}function w(){e.items=[],e.count=-1}async function D(){const{data:n}=await c.get(u);return n}const m={configs:l(e)};function S(){return m}export{s as C,r as a,T as b,g as c,y as d,k as e,D as f,w as g,b as h,S as u};
