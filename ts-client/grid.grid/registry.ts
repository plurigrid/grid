import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgSubmitEstimate } from "./types/grid/grid/tx";
import { MsgSubmitPower } from "./types/grid/grid/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/grid.grid.MsgSubmitEstimate", MsgSubmitEstimate],
    ["/grid.grid.MsgSubmitPower", MsgSubmitPower],
    
];

export { msgTypes }