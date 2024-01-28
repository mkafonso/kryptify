import { z } from "zod";

export const createAccountSchema = z.object({
  name: z.string().toLowerCase().nonempty("name cannot be empty"),
  email: z.string().toLowerCase().trim().email("invalid email address"),
  password: z.string().min(8, "password must be at 8 least characterts long"),
});
