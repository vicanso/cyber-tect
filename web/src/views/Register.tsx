import { defineComponent } from "vue";
import ExLoginRegister from "../components/ExLoginRegister";

export default defineComponent({
  name: "Register",
  render() {
    return <ExLoginRegister type="register" />;
  },
});
