import { generateRoutersTs } from "./generateRoutersTs";
import { generateMainTs } from "./generateMainTs";

export const generateCodes = () => {
  generateRoutersTs();
  generateMainTs();
};
