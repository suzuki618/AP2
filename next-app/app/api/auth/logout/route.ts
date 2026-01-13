// app/api/auth/logout/route.ts
import { NextRequest, NextResponse } from 'next/server';
import { SupabaseAuthService } from '@/lib/supabaseAuthService';

export async function POST(req: NextRequest) {
  // Authorization ヘッダ取得
  const authHeader = req.headers.get('authorization');

  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return NextResponse.json(
      { error: 'Unauthorized' },
      { status: 401 }
    );
  }

  const accessToken = authHeader.slice(7); // "Bearer " を除去

  // Supabaseへログアウト要求
  const { json, status } =
    await SupabaseAuthService.logout(accessToken);

  return NextResponse.json(json, { status });
}
