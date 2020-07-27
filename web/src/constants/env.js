const env = process.env.NODE_ENV;

export function isDevelopment() {
  return env === "development";
}

export function isProduction() {
  return env === "production";
}
