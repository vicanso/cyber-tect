import { names } from "./routes";
import router from "./router";
import { LocationQueryRaw } from "vue-router";

export function goTo(
  name: string,
  params: {
    replace?: boolean;
    query?: LocationQueryRaw;
  }
): void {
  router.push({
    name,
    replace: params.replace,
    query: params.query,
  });
}

export function goToLogin(replace?: boolean): void {
  goTo(names.login, {
    replace: replace ?? false,
  });
}

export function goToRegister(replace?: boolean): void {
  goTo(names.register, {
    replace: replace ?? false,
  });
}

export function goToHome(replace?: boolean): void {
  goTo(names.home, {
    replace: replace ?? false,
  });
}

export function goToProfile(replace?: boolean): void {
  goTo(names.profile, {
    replace: replace ?? false,
  });
}
