import * as outscale from "@pulumi/outscale";

const random = new outscale.Random("my-random", { length: 24 });

export const output = random.result;