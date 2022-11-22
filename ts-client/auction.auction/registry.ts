import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgFinalizeAuction } from "./types/auction/auction/tx";
import { MsgPlaceBid } from "./types/auction/auction/tx";
import { MsgCreateAuction } from "./types/auction/auction/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/auction.auction.MsgFinalizeAuction", MsgFinalizeAuction],
    ["/auction.auction.MsgPlaceBid", MsgPlaceBid],
    ["/auction.auction.MsgCreateAuction", MsgCreateAuction],
    
];

export { msgTypes }