/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as ExampleApiCommonP from "./common_p.pb";
import * as fm from "./fetch.pb";


export type VersionRequest = {
  emp?: ExampleApiCommonP.Empty;
};

export type VersionResponse = {
  version?: string;
};

export type Version = Record<string, never>;

export class VelezAPI {
  static Version(this:void, req: VersionRequest, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchRequest<VersionResponse>(`/api/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"});
  }
}