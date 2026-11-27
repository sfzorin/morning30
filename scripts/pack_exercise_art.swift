#!/usr/bin/env swift
// macOS: swiftc -O scripts/pack_exercise_art.swift -o /tmp/pack_exercise_art
// /tmp/pack_exercise_art SOURCE.png FRAMES still|movement|sequence|alternating OUTPUT_DIR [--registration FILE]
// The output directory contains only this character's media for one exercise.
import CoreGraphics
import Darwin
import Foundation
import ImageIO
import UniformTypeIdentifiers
import zlib

struct PackError: Error, CustomStringConvertible {
    let description: String
    init(_ message: String) { description = message }
}

let side = 512
let rgb = CGColorSpace(name: CGColorSpace.sRGB)!
let bitmapFormat = CGImageAlphaInfo.noneSkipLast.rawValue | CGBitmapInfo.byteOrder32Big.rawValue
let files = FileManager.default
let mediaNames = Set(["still.webp", "movement.webp", "sequence.webp", "alternating.webp"])

func readPNG(_ input: URL) throws -> CGImage {
    let size = try input.resourceValues(forKeys: [.fileSizeKey]).fileSize ?? 0
    guard (45...64*1024*1024).contains(size) else { throw PackError("Invalid PNG file size") }
    let data = try Data(contentsOf: input, options: .mappedIfSafe)
    // ImageIO can display partially damaged images. Check chunk lengths and CRCs before decoding.
    try data.withUnsafeBytes { raw in
        let bytes = raw.bindMemory(to: UInt8.self)
        func uint32(_ i: Int) -> Int {
            (Int(bytes[i]) << 24) | (Int(bytes[i+1]) << 16) | (Int(bytes[i+2]) << 8) | Int(bytes[i+3])
        }
        guard bytes.starts(with: [137, 80, 78, 71, 13, 10, 26, 10]) else { throw PackError("Expected PNG input") }
        var offset = 8
        while offset+12 <= bytes.count {
            let length = uint32(offset)
            guard length <= bytes.count-offset-12 else { throw PackError("Truncated PNG chunk") }
            let tag = String(decoding: bytes[offset+4..<offset+8], as: UTF8.self)
            guard crc32(0, bytes.baseAddress!+offset+4, UInt32(length+4)) == uint32(offset+8+length)
            else { throw PackError("Invalid PNG checksum in \(tag)") }
            if offset == 8 {
                guard tag == "IHDR", length == 13 else { throw PackError("Missing PNG header") }
                let w = uint32(offset+8), h = uint32(offset+12)
                guard w > 0, h > 0, w <= 16384, h <= 16384, w*h <= 40_000_000
                else { throw PackError("Unsupported PNG dimensions \(w)×\(h)") }
            }
            guard tag != "acTL" else { throw PackError("Expected one sheet, received animated PNG") }
            offset += length+12
            if tag == "IEND" {
                guard length == 0, offset == bytes.count else { throw PackError("Invalid PNG ending") }
                return
            }
        }
        throw PackError("Missing PNG end chunk")
    }
    guard let src = CGImageSourceCreateWithData(data as CFData, nil), CGImageSourceGetCount(src) == 1,
          let image = CGImageSourceCreateImageAtIndex(src, 0, [kCGImageSourceShouldCacheImmediately: true] as CFDictionary),
          CGImageSourceGetStatusAtIndex(src, 0) == .statusComplete else { throw PackError("Could not decode complete PNG") }
    return image
}

func canvas(_ width: Int, _ height: Int) throws -> CGContext {
    guard let c = CGContext(data: nil, width: width, height: height, bitsPerComponent: 8,
                            bytesPerRow: width*4, space: rgb, bitmapInfo: bitmapFormat)
    else { throw PackError("Could not allocate image canvas") }
    c.setFillColor(CGColor(gray: 1, alpha: 1))
    c.fill(CGRect(x: 0, y: 0, width: width, height: height))
    c.interpolationQuality = .high
    return c
}

func checkLayout(_ image: CGImage, _ frames: Int) throws {
    let cellWidth = Double(image.width)/Double(frames)
    let ratio = cellWidth/Double(image.height)
    guard cellWidth >= 64, image.height >= 64, (0.24...4.0).contains(ratio)
    else { throw PackError("Implausible dimensions for \(frames) horizontal frame(s): \(image.width)×\(image.height)") }

    // Equal-size analysis thumbnails only. Leave source pixels and equal column cuts intact.
    let tile = 96, width = frames*tile
    let preview = try canvas(width, tile)
    preview.draw(image, in: CGRect(x: 0, y: 0, width: width, height: tile))
    let p = preview.data!.assumingMemoryBound(to: UInt8.self)
    func ink(_ x: Int, _ y: Int) -> Bool {
        let i = y*preview.bytesPerRow+x*4
        return min(p[i], p[i+1], p[i+2]) < 175
    }
    var rows = Array(repeating: Array(repeating: 0, count: tile), count: frames)
    for column in 0..<frames {
        let start = column*tile
        var inkCount = 0, crossing = 0
        for y in 0..<tile {
            for x in start..<start+tile where ink(x,y) { rows[column][y] += 1 }
            inkCount += rows[column][y]
            if start > 0 && ink(start-1,y) && ink(start,y) { crossing += 1 }
        }
        guard inkCount >= 64 else { throw PackError("Frame \(column+1) is empty or nearly blank") }
        guard crossing < tile/5 else { throw PackError("Column boundary cuts through artwork before frame \(column+1)") }
    }
    // Two substantial subjects separated by the same empty band in every column indicate stacked rows.
    var gap = 0
    for y in 0..<tile where frames > 1 {
        gap = rows.allSatisfy { $0[y] < 2 } ? gap+1 : 0
        if gap == tile/8, y < tile-1 {
            let stacked = rows.allSatisfy { row in
                let total = row.reduce(0,+)
                let above = row.prefix(y+1-gap).reduce(0,+)
                let below = row.suffix(from: y+1).reduce(0,+)
                return min(above,below)*4 > total
            }
            if stacked { throw PackError("Expected a single horizontal row, found two rows of artwork") }
        }
    }
}

// Manual output-pixel offsets preserve intentional motion. Registration is opt-in per exercise.
struct Registration: Decodable {
    let exercise: String
    let offsets: [String: [[Int]]]
}

func readOffsets(_ file: URL, target: URL, frames: Int) throws -> [CGPoint] {
    let recipe = try JSONDecoder().decode(Registration.self, from: Data(contentsOf: file))
    guard target.deletingLastPathComponent().lastPathComponent == recipe.exercise,
          let offsets = recipe.offsets[target.lastPathComponent], offsets.count == frames,
          offsets.allSatisfy({ $0.count == 2 && $0.allSatisfy({ (-side..<side).contains($0) }) })
    else { throw PackError("Registration must match the exercise, character and frame count") }
    return offsets.map { CGPoint(x: $0[0], y: $0[1]) }
}

func makeStrip(_ input: CGImage, offsets: [CGPoint]) throws -> CGImage {
    let frames = offsets.count
    let strip = try canvas(side*frames, side)
    for i in 0..<frames {
        // Rounding the boundaries includes every source pixel for widths such as 1024 / 3.
        let left = (input.width*i+frames/2)/frames
        let right = (input.width*(i+1)+frames/2)/frames
        let rect = CGRect(x: left, y: 0, width: right-left, height: input.height)
        guard let cell = input.cropping(to: rect) else { throw PackError("Could not read frame \(i+1)") }
        let scale = CGFloat(side)/max(rect.width, rect.height)
        let w = rect.width*scale, h = rect.height*scale
        let tile = CGRect(x: i*side, y: 0, width: side, height: side)
        let destination = CGRect(x: tile.midX-w/2+offsets[i].x, y: tile.midY-h/2-offsets[i].y, width: w, height: h)
        strip.saveGState()
        strip.clip(to: tile)
        strip.draw(cell, in: destination)
        strip.restoreGState()
    }
    guard let image = strip.makeImage() else { throw PackError("Could not create strip") }
    return image
}

// Catalog covers use the most recognizable phase. Standalone packs retain the first.
func coverOffset(_ target: URL, frames: Int) throws -> Int {
    guard frames > 1 else { return 0 }
    let exercise = target.deletingLastPathComponent()
    let manifest = exercise.deletingLastPathComponent().appendingPathComponent("motions.json")
    guard files.fileExists(atPath: manifest.path) else { return 0 }
    struct Motion: Decodable { let cover: Int? }
    let entries = try JSONDecoder().decode([String: Motion].self, from: Data(contentsOf: manifest))
    let number = entries[exercise.lastPathComponent]?.cover ?? 1
    guard (1...frames).contains(number) else { throw PackError("Cover phase out of range for \(exercise.lastPathComponent)") }
    return (number - 1) * side
}

func writeWebP(_ image: CGImage, to output: URL, encoder: URL) throws {
    let png = output.deletingPathExtension().appendingPathExtension("png")
    defer { try? files.removeItem(at: png) }
    guard let destination = CGImageDestinationCreateWithURL(png as CFURL, UTType.png.identifier as CFString, 1, nil)
    else { throw PackError("Could not stage PNG") }
    CGImageDestinationAddImage(destination, image, nil)
    guard CGImageDestinationFinalize(destination) else { throw PackError("Could not write staged PNG") }
    let process = Process()
    process.executableURL = encoder
    process.arguments = ["-quiet", "-q", "87", "-m", "5", "-sharp_yuv", png.path, "-o", output.path]
    try process.run()
    process.waitUntilExit()
    guard process.terminationStatus == 0 else { throw PackError("cwebp failed for \(output.lastPathComponent)") }
}

func publish(_ staged: URL, to target: URL) throws {
    if renamex_np(staged.path, target.path, UInt32(RENAME_EXCL)) == 0 { return }
    guard errno == EEXIST else { throw PackError("Could not publish \(target.path): \(String(cString: strerror(errno)))") }
    let info = try target.resourceValues(forKeys: [.isDirectoryKey, .isSymbolicLinkKey])
    guard info.isDirectory == true, info.isSymbolicLink != true else { throw PackError("Target must be a real media directory") }
    for entry in try files.contentsOfDirectory(at: target, includingPropertiesForKeys: [.isRegularFileKey, .isSymbolicLinkKey]) {
        let meta = try entry.resourceValues(forKeys: [.isRegularFileKey, .isSymbolicLinkKey])
        guard mediaNames.contains(entry.lastPathComponent), meta.isRegularFile == true, meta.isSymbolicLink != true
        else { throw PackError("Refusing to replace unrelated content: \(entry.path)") }
    }
    // Swap both WebPs as a unit. A reader (or concurrent writer) only sees a complete set.
    guard renamex_np(staged.path, target.path, UInt32(RENAME_SWAP)) == 0
    else { throw PackError("Could not replace \(target.path): \(String(cString: strerror(errno)))") }
}

func run() throws {
    let args = CommandLine.arguments
    guard (args.count == 5 || args.count == 7 && args[5] == "--registration"), let frames = Int(args[2]),
          (frames == 1 && args[3] == "still" || [2,3].contains(frames) && ["movement","sequence"].contains(args[3]) ||
            frames == 3 && args[3] == "alternating") else {
        throw PackError("Usage: pack_exercise_art SOURCE.png FRAMES KIND OUTPUT_DIR [--registration FILE] (1 still, 2–3 movement/sequence, 3 alternating)")
    }
    let input = URL(fileURLWithPath: args[1]), target = URL(fileURLWithPath: args[4]).standardizedFileURL
    let paths = (ProcessInfo.processInfo.environment["PATH"] ?? "").split(separator: ":")
    guard let encoder = paths.map({ URL(fileURLWithPath: String($0)).appendingPathComponent("cwebp") }).first(where: { files.isExecutableFile(atPath: $0.path) })
    else { throw PackError("cwebp is required on PATH") }
    let image = try readPNG(input)
    try checkLayout(image, frames)
    let offsets = args.count == 7 ? try readOffsets(URL(fileURLWithPath: args[6]), target: target, frames: frames)
                                  : Array(repeating: CGPoint.zero, count: frames)
    let strip = try makeStrip(image, offsets: offsets)
    let coverX = try coverOffset(target, frames: frames)
    guard let still = strip.cropping(to: CGRect(x: coverX, y: 0, width: side, height: side)) else { throw PackError("Could not extract still") }
    let parent = target.deletingLastPathComponent()
    try files.createDirectory(at: parent, withIntermediateDirectories: true)
    let staged = parent.appendingPathComponent(".art-"+UUID().uuidString, isDirectory: true)
    try files.createDirectory(at: staged, withIntermediateDirectories: false)
    defer { try? files.removeItem(at: staged) }
    try writeWebP(still, to: staged.appendingPathComponent("still.webp"), encoder: encoder)
    if frames > 1 { try writeWebP(strip, to: staged.appendingPathComponent(args[3]+".webp"), encoder: encoder) }
    try publish(staged, to: target)
    print("\(image.width)×\(image.height) → \(frames) × \(side)×\(side)   \(target.path)")
}

do { try run() } catch {
    FileHandle.standardError.write(Data("pack_exercise_art: \(error)\n".utf8))
    exit(1)
}
